package domain

import (
	"gorm.io/gorm"
	"time"
)

type PlantHireRequestStatus string

const (
	PhrStatusAccepted PlantHireRequestStatus = "accepted"
	PhrStatusRejected PlantHireRequestStatus = "rejected"
	PhrStatusPending  PlantHireRequestStatus = "pending"
)

type PurchaseOrderStatus string

const (
	PoStatusAccepted PurchaseOrderStatus = "accepted"
	PoStatusRejected PurchaseOrderStatus = "rejected"
)

type PlantHireRequest struct {
	gorm.Model
	ID                    int
	SiteEngineerName      string
	ConstructionSiteName  string
	Comment               string
	Cost                  float64
	PlantInventoryEntryID int
	PurchaseOrderStatus   PurchaseOrderStatus
	StartDate             time.Time
	EndDate               time.Time
	Status                PlantHireRequestStatus
	CreatedAt             time.Time
}
