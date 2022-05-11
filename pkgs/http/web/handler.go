package web

import (
	"net/http"
	"strconv"
	"time"

	models "github.com/devbyP/webboard/pkgs/storage/db-gorm"
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

func (hp *homePost) getPost(p *models.Post) {
	hp.ID = p.ID
	hp.Content = p.Content
	hp.EditedAt = p.UpdatedAt
	hp.PostAt = p.CreatedAt
	hp.ShortDesc = p.Content
	hp.Title = p.Title
}

type homeContext struct {
	Posts []*homePost
}

func (hc *homeContext) appendPost(post *homePost) {
	hc.Posts = append(hc.Posts, post)
}

var postPerPage int = 10

func getCurrentPage(c echo.Context) int {
	pageNum, err := getParamIDNum(c, "page")
	if err != nil {
		return 1
	}
	return pageNum
}

func GetPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		page := getCurrentPage(c)
		limit := postPerPage
		offset := (page - 1) * limit

		post := &models.Post{}
		hc := &homeContext{}

		posts, err := post.GetAllPosts(db, limit, offset)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error Getting post from server.")
		}

		for _, ps := range posts {
			hp := &homePost{}
			hp.getPost(ps)
			hc.appendPost(hp)
		}
		return c.Render(http.StatusOK, "home.tmpl", hc)
	}
}
