package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/devbyP/webboard/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getParamIDNum(c echo.Context, param string) (int, error) {
	return strconv.Atoi(c.Param(param))
}

type postContext struct {
	ID       uint
	Title    string
	Content  string
	PostAt   time.Time
	EditedAt time.Time
}

func GetPost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := getParamIDNum(c, "id")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Required id param as number type")
		}
		newPost := &models.Post{}
		db.First(newPost, id)

		return c.Render(http.StatusOK, "post.tmpl", postContext{
			ID:       newPost.ID,
			Title:    newPost.Title,
			Content:  newPost.Content,
			PostAt:   newPost.CreatedAt,
			EditedAt: newPost.UpdatedAt,
		})
	}
}

type homePost struct {
	ID        uint
	Title     string
	Content   string
	ShortDesc string
	PostAt    time.Time
	EditedAt  time.Time
}

type homeContext struct {
	Posts []homePost
}

func (hc *homeContext) appendPost(post homePost) {
	hc.Posts = append(hc.Posts, post)
}

func GetPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := getParamIDNum(c, "page")
		if err != nil || page <= 0 {
			page = 1
		}
		limit := 10
		offset := (page - 1) * limit
		rows, err := db.Model(&models.Post{}).Order("created_at").Limit(limit).Offset(offset).Rows()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error Getting post from server.")
		}

		for rows.Next() {
			post := models.Post{}
			db.ScanRows(rows, &post)
		}
		return c.Render(http.StatusOK, "home.tmpl", homeContext{})
	}
}
