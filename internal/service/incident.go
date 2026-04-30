package service

import (
	"errors"
	"time"

	"pulseops/internal/model"
	"pulseops/internal/repository"
)

type IncidentService struct {
	repo *repository.IncidentRepository
}

func NewIncidentService(repo *repository.IncidentRepository) *IncidentService {
	return &IncidentService{repo: repo}
}

func (s *IncidentService) GetAll() ([]model.Incident, error) {
	return s.repo.FindAll()
}

func (s *IncidentService) GetByID(id uint) (model.Incident, error) {
	incident, err := s.repo.FindByID(id)
	if err != nil {
		return model.Incident{}, errors.New("incident not found")
	}
	return incident, nil
}

func (s *IncidentService) Create(incident *model.Incident) error {
	if err := validateIncidentDates(incident); err != nil {
		return err
	}
	return s.repo.Create(incident)
}

func (s *IncidentService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("incident not found")
	}
	return s.repo.Delete(id)
}

func validateIncidentDates(i *model.Incident) error {
	now := time.Now().UTC()

	if i.StartedAt.Before(now.AddDate(0, 0, -30)) {
		return errors.New("started_at cannot be more than 30 days in the past")
	}

	if i.StartedAt.After(now) {
		return errors.New("started_at cannot be in the future")
	}

	return nil
}
