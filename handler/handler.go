package handler

import "github.com/SBEKzy/testTask/service"

type handler struct {
	Service service.Service
}

func New(s service.Service) *handler {
	return &handler{
		Service: s,
	}
}
