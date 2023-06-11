package services

import (
	"backend/models/migrations"
	"backend/repositories"
)

type RefJafungService interface {
	FindAll() ([]migrations.RefJabFung, error)
}

type service struct {
	repository repositories.RefJafungRepository
}

func NewRefJafungService(repository repositories.RefJafungRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]migrations.RefJabFung, error) {
	return s.repository.FindAll()
}
