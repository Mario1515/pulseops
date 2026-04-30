package repository

import (
	"pulseops/internal/model"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) FindAll() ([]model.Incident, error) {
	var incidents []model.Incident
	return incidents, r.db.Order("started_at DESC").Find(&incidents).Error
}

func (r *IncidentRepository) FindByID(id uint) (model.Incident, error) {
	var incident model.Incident
	return incident, r.db.First(&incident, id).Error
}

func (r *IncidentRepository) Create(incident *model.Incident) error {
	return r.db.Create(incident).Error
}

func (r *IncidentRepository) Update(incident *model.Incident) error {
	return r.db.Save(incident).Error
}

func (r *IncidentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Incident{}, id).Error
}
