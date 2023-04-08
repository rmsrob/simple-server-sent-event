package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

/**
 * We are computing a message to be passed into the channel
 * to be read and be sent into the openned connection on /event
**/
func GetPosted(msgChan chan string) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		// Send a message to the msgChan channel with the current time formatted in European time
		msg := fmt.Sprintf("%s - Hello, world!", time.Now().Format("02-01-2006 15:04:05 MST"))
		msgChan <- msg

		return nil
	}
}
