package main

import (
	"github.com/bryutus/caspian-back/app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost"},
		AllowMethods: []string{echo.GET},
	}))

	e.GET("/album", handler.GetAlbum())
	e.GET("/song", handler.GetSong())

	e.Start(":1323")
}
