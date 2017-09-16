package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAlbum() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")

		j := map[string]string{
			"resource": "album",
			"limit":    limit,
		}

		return c.JSON(http.StatusOK, j)
	}
}

func GetSong() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")

		j := map[string]string{
			"resource": "song",
			"limit":    limit,
		}

		return c.JSON(http.StatusOK, j)
	}
}