package service

import (
	"github.com/SBEKzy/testTask/model"
	"github.com/SBEKzy/testTask/repository"
)

type Service interface {
	SendRequest(requestModel *model.Request) (*model.Response, error)
}

type service struct {
	Repo repository.Repository
}

func New(r repository.Repository) Service {
	return &service{
		Repo: r,
	}
}
