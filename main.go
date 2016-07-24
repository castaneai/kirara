package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	kirara "github.com/castaneai/kirara/api"
)

func main() {
	e := echo.New()
	api := e.Group("/api/v1")
	{
		api.Get("/posts/:keyword", kirara.GetPosts)
	}
	e.Run(standard.New("127.0.0.1:8888"))
}
