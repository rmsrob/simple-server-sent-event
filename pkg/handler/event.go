package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * Open the SSE connection and return the events from post-events endpoint
 * when they happened
**/
func Event(msgChan chan string) echo.HandlerFunc {
	return func(c echo.Context) error {
		const MIMETextEventStream = "text/event-stream"
		c.Response().Header().Set(echo.HeaderContentType, MIMETextEventStream)
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		// Start an infinite loop to keep listening to the msg channel for new messages
		for {
			select {
			case msg := <-msgChan:
				fmt.Println("Client connected")
				fmt.Printf("data: %s\n\n", msg)
				c.Response().Flush()
			case <-c.Request().Context().Done():
				fmt.Println("Client closed connection")
				return nil
			}
		}
	}
}