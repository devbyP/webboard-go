package web

import (
	"net/http"
	"time"

	models "github.com/devbyP/webboard/pkgs/storage/db-gorm"
	"github.com/labstack/echo/v4"
)

type Post struct {
	ID        int
	Title     string
	DescShort string
	PostDate  time.Time
}

func getPost() []*Post {
	posts := make([]*Post, 0)
	post1 := &Post{
		ID:        1,
		Title:     "Hi I am first post.",
		DescShort: "Hello world first one baby.",
		PostDate:  time.Now().Add(time.Duration(-20) * time.Minute),
	}
	post2 := &Post{
		ID:        2,
		Title:     "Yo I'm in",
		DescShort: "So excited to be here this is my first post.",
		PostDate:  time.Now().Add(time.Duration(-10) * time.Minute),
	}
	post3 := &Post{
		ID:        3,
		Title:     "How you guys so fast?",
		DescShort: "I'm rushing to signup to this web site and though i'm the first one.",
		PostDate:  time.Now(),
	}
	posts = append(posts, post1, post2, post3)
	return posts
}

func URLMapping(e *echo.Echo) {
	res := struct {
		Test string
		Post []*Post
	}{
		Test: "test message",
		Post: getPost(),
	}
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.tmpl", res)
	})
	e.GET("/profile/:name", func(c echo.Context) error {
		name := c.Param("name")
		user := struct {
			Username string
		}{
			Username: name,
		}
		return c.Render(http.StatusOK, "profile.tmpl", user)
	})
	e.GET("/post/:id", GetPost(models.GetGorm()))
}
