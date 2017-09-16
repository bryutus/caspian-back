package main

import (
	"github.com/bryutus/caspian-back/app/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/album", handler.GetAlbum())
	e.GET("/song", handler.GetSong())

	e.Start(":1323")
}
