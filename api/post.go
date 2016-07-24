package api

import (
	"github.com/castaneai/kirara/resource"
	"github.com/labstack/echo"
	"net/http"
)

func GetPosts(c echo.Context) error {
	keyword := c.Param("keyword")
	res := resource.Tiqav{} // TODO: other resources
	posts, err := res.SearchByKeyword(keyword, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, posts)
}