package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func customHttpErrHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("error%d.tmpl", code)

	if renErr := c.Render(code, errorPage, err.Error()); renErr != nil {
		c.Logger().Error(renErr)
		c.Render(code, "error.tmpl", err.Error())
	}
}

type HandlerMapingFunc func(*echo.Echo)

func StartServer(maping HandlerMapingFunc) {
	e := echo.New()
	if maping == nil {
		log.Fatal("error, no url mapping found")
	}
	e.Static("/", "webContent")
	temp := template.Must(template.ParseGlob("templates/*.tmpl"))
	//temp = template.Must(temp.ParseGlob("templates/error/*.tmpl"))
	//temp = template.Must(temp.ParseGlob("templates/includes/*.tmpl"))
	t := &Template{
		Template: temp,
	}
	e.Renderer = t
	e.HTTPErrorHandler = customHttpErrHandler
	e.Use(middleware.Logger())
	maping(e)
	e.Logger.Fatal(e.Start(":" + getPort()))
}

func getPort() (port string) {
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}
	return
}
