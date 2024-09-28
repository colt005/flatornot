package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleSSE(e echo.Context) error {
	e.Response().Header().Set("Content-Type", "text/event-stream")
	e.Response().Header().Set("Cache-Control", "no-cache")
	e.Response().Header().Set("Connection", "keep-alive")

	client := make(chan string)

	h.service.RegisterClient(client)

	defer func() {
		h.service.UnRegisterClient(client)
	}()

	go func() {
		<-e.Request().Context().Done()
	}()

	for {
		select {
		case <-e.Request().Context().Done():
			return nil
		case data := <-client:
			fmt.Fprintf(e.Response().Writer, "data: %s\n\n", data)
			e.Response().Flush()
		}
	}

}

func (h *Handler) HandleBroadcast() {
	h.service.HandleBroadcast()
}
