package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/bryutus/caspian-serverside/app/db"
	"github.com/bryutus/caspian-serverside/app/models"
	"github.com/labstack/echo"
)

func GetAlbums() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		if matched, _ := regexp.MatchString(`[0-9]`, limit); !matched {
			return c.JSON(http.StatusOK, fmt.Sprintf("Invalid value %s", limit))
		}

		db := db.Connect()
		defer db.Close()

		h := models.History{}
		if db.Where("resource_type = ?", "album").Last(&h).RecordNotFound() {
			return c.JSON(http.StatusOK, "not found")
		}

		r := []models.Resource{}
		if db.Model(&h).Limit(limit).Related(&r).RecordNotFound() {
			return c.JSON(http.StatusOK, "not found")
		}

		return c.JSON(http.StatusOK, r)
	}
}

func GetSongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")

		j := map[string]string{
			"resource": "song",
			"limit":    limit,
		}

		return c.JSON(http.StatusOK, j)
	}
}
