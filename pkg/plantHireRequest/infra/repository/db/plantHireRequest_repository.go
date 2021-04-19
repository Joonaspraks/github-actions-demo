package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/alexgtn/buildit/pkg/domain"
	"gorm.io/gorm"
)

const plantHireRequestTable = "plant_hire_request"

// PlantHireRequestRepository is a SQL implementation of a phr repository
type PlantHireRequestRepository struct {
	db *gorm.DB
}

// NewPlantHireRequestRepository builds new SQL phr repository
func NewPlantHireRequestRepository(db *gorm.DB) *PlantHireRequestRepository {
	return &PlantHireRequestRepository{
		db: db,
	}
}

// Create creates PHR
func (r *PlantHireRequestRepository) Create(phr *domain.PlantHireRequest) (*domain.PlantHireRequest, error) {
	tx := r.db.Create(phr)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, fmt.Errorf("error inserting plant hire request %v", phr)
	}
	return phr, nil
}

// GetAll gets all PHRs
func (r *PlantHireRequestRepository) GetAll() ([]*domain.PlantHireRequest, error) {
	var phrs []*domain.PlantHireRequest
	tx := r.db.Find(&phrs)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, fmt.Errorf("error querying phrs FROM table %s", plantHireRequestTable)
	}

	return phrs, nil
}
