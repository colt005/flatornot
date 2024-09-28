package handlers

import "github.com/colt005/flatornot/service"

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}


