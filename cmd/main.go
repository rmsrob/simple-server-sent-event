package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rrobrms/simple-server-sent-event/pkg/handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the msgChan channel
	msgChan := make(chan string)

	// Routes
	e.GET("/event", handler.Event(msgChan))
	e.GET("/post-event", handler.GetPosted(msgChan))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
