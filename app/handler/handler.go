package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAlbum() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		return c.String(http.StatusOK, "album limit"+limit)
	}
}

func GetSong() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		return c.String(http.StatusOK, "song limit"+limit)
	}
}
