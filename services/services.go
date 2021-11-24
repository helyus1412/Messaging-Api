package services

import (
	"github.com/helyus1412/messaging-api/entities"
	"github.com/helyus1412/messaging-api/repository"
)

type Services interface {
	AllMessages() ([]entities.Messages, error)
}

type services struct {
	repo repository.Repository
}

func NewServices(repo repository.Repository) *services {
	return &services{repo}
}

func (s *services) AllMessages() ([]entities.Messages, error) {
	allMessages, err := s.repo.AllMessages()

	if err != nil {
		return allMessages, err
	}

	return allMessages, nil
}
