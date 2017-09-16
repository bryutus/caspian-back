package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
)

func GetAlbums() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		if matched, _ := regexp.MatchString(`[0-9]`, limit); !matched {
			return c.JSON(http.StatusOK, fmt.Sprintf("Invalid value %s", limit))
		}

		return c.JSON(http.StatusOK, j)
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
