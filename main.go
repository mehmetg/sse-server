package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const(
	MIMEApplicationEventStream = "text/event-stream"
	HeaderCacheControl = "Cache-Control"
	HeaderConnection = "Connection"
	ConnectionKeepAlive = "keep-alive"
	CacheNoCache = "no-cache"
)

type (
	Event struct {
		Id string `json:"id"`
		Data string `json:"data"`
	}
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, MIMEApplicationEventStream)
		c.Response().Header().Set(HeaderCacheControl, CacheNoCache)
		c.Response().Header().Set(HeaderConnection, ConnectionKeepAlive)
		c.Response().WriteHeader(http.StatusOK)
		for {
			responseID := uuid.New().String()
			responseData := uuid.New().String()
			event := Event{responseID, responseData}
			if err := json.NewEncoder(c.Response()).Encode(event); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(3 * time.Second)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
