package server

import (
	"log/slog"
	"os"

	"github.com/colt005/flatornot/handlers"
	"github.com/colt005/flatornot/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	handler *handlers.Handler
	e       *echo.Echo
}

func New() (*Server, error) {
	s, err := service.New()
	if err != nil {
		return nil, err
	}
	h := handlers.New(s)
	ser := &Server{
		handler: h,
		e:       echo.New(),
	}

	//ser.e.Use(middleware.Logger())
	ser.e.Use(middleware.Recover())

	return ser, nil
}

func (s *Server) RegisterRoutes() {
	s.e.GET("/", s.handler.HandleIndex)
	s.e.POST("/vote", s.handler.HandleVote)
	s.e.Any("/events", s.handler.HandleSSE)

	go s.handler.HandleBroadcast()

}

func (s *Server) Start() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	listenAddr = ":3000"
	slog.Info("Server Started", "ListenAddress", listenAddr)
	s.e.Logger.Fatal(s.e.Start(listenAddr))
}

