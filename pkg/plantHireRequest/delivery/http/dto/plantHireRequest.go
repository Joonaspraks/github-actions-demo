package dto

import (
	"errors"
	"gitlab.com/alexgtn/buildit/pkg/domain"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/common/validation"
	"time"
)

var errPlantHireRequestNil = errors.New("error plant hire request is nil")

type PlantHireRequest struct {
	ID                    int        `json:"id"`
	SiteEngineerName      string     `json:"siteEngineerName"`
	ConstructionSiteName  string     `json:"constructionSiteName"`
	Comment               string     `json:"comment"`
	Cost                  float64    `json:"cost"`
	PlantInventoryEntryID int        `json:"plantInventoryEntryID"`
	PurchaseOrderStatus   string     `json:"purchaseOrderStatus"`
	StartDate             *time.Time `json:"startDate"`
	EndDate               *time.Time `json:"endDate"`
	Status                string     `json:"status"`
	CreatedAt             *time.Time `json:"createdAt"`
}

func NewPlantHireRequest(phr *domain.PlantHireRequest) (*PlantHireRequest, error) {
	if phr == nil {
		return nil, errPlantHireRequestNil
	}

	return &PlantHireRequest{
		ID:                    phr.ID,
		SiteEngineerName:      phr.SiteEngineerName,
		ConstructionSiteName:  phr.ConstructionSiteName,
		Comment:               phr.Comment,
		Cost:                  phr.Cost,
		PlantInventoryEntryID: phr.PlantInventoryEntryID,
		PurchaseOrderStatus:   string(phr.PurchaseOrderStatus),
		StartDate:             &phr.StartDate,
		EndDate:               &phr.EndDate,
		Status:                string(phr.Status),
		CreatedAt:             &phr.CreatedAt,
	}, nil
}

// News builds new list of PHR DTOs
func NewPlantHireRequests(phrs []*domain.PlantHireRequest) (res []*PlantHireRequest, err error) {
	for _, item := range phrs {
		it, err := NewPlantHireRequest(item)
		if err != nil {
			return nil, err
		}
		res = append(res, it)
	}
	return res, nil
}

// Validate is used to validate the body
func (r *PlantHireRequest) Validate() error {
	errs := validation.NewErrorGroup()

	if r.SiteEngineerName == "" {
		errs.AddError(validation.New("siteEngineerName", "siteEngineerName cannot be empty"))
	}

	if r.ConstructionSiteName == "" {
		errs.AddError(validation.New("constructionSiteName", "constructionSiteName cannot be empty"))
	}

	if r.StartDate == nil {
		errs.AddError(validation.New("startDate", "startDate cannot be empty"))
	}

	if r.EndDate == nil {
		errs.AddError(validation.New("endDate", "endDate cannot be empty"))
	}

	if r.Cost <= 0 {
		errs.AddError(validation.New("cost", "cost cannot be zero or negative"))
	}

	if r.PlantInventoryEntryID <= 0 {
		errs.AddError(validation.New("plantInventoryEntryID", "plantInventoryEntryID invalid"))
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}
