package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * Open the SSE connection and return the events from /event endpoint when they happened
**/
func Listener(msgChan chan string) echo.HandlerFunc {
	return func(c echo.Context) error {
		const MIMETextEventStream = "text/event-stream"
		c.Response().Header().Set(echo.HeaderContentType, MIMETextEventStream)
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		for {
			select {
			case msg := <-msgChan:
				fmt.Printf("event: %s\n\n", msg)
				c.Response().Flush()
			case <-c.Request().Context().Done():
				fmt.Println("Client closed connection")
				return nil
			}
		}
	}
}
