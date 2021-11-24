package repository

import (
	"github.com/helyus1412/messaging-api/entities"
	"gorm.io/gorm"
)

type Repository interface {
	AllMessages() ([]entities.Messages, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AllMessages() ([]entities.Messages, error) {
	var allMessages []entities.Messages

	if err := r.db.Find(&allMessages).Error; err != nil {
		return allMessages, err
	}

	return allMessages, nil
}
