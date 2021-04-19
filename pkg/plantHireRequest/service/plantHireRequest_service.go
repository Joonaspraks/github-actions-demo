package service

import (
	"gitlab.com/alexgtn/buildit/pkg/domain"
	"time"
)

type plantHireRequestRepository interface {
	Create(phr *domain.PlantHireRequest) (*domain.PlantHireRequest, error)
	GetAll() ([]*domain.PlantHireRequest, error)
}

// PlantHireRequestService service
type PlantHireRequestService struct {
	phrRepository plantHireRequestRepository
}

// NewPlantHireRequestService
func NewPlantHireRequestService(phrRepository plantHireRequestRepository) *PlantHireRequestService {
	return &PlantHireRequestService{
		phrRepository: phrRepository,
	}
}

// Create creates a PHR
func (uc *PlantHireRequestService) Create(siteEngineerName string, constructionSiteName string, comment string, cost float64, plantInventoryEntryID int, startDate *time.Time, endDate *time.Time) (*domain.PlantHireRequest, error) {
	phr := &domain.PlantHireRequest{
		SiteEngineerName:      siteEngineerName,
		ConstructionSiteName:  constructionSiteName,
		Comment:               comment,
		Cost:                  cost,
		PlantInventoryEntryID: plantInventoryEntryID,
		StartDate:             *startDate,
		EndDate:               *endDate,
		Status:                domain.PhrStatusPending,
	}

	return uc.phrRepository.Create(phr)
}

// GetAll gets all PHRs
func (uc *PlantHireRequestService) GetAll() ([]*domain.PlantHireRequest, error) {
	return uc.phrRepository.GetAll()
}
