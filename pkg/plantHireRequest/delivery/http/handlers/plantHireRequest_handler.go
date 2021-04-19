package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/alexgtn/buildit/pkg/domain"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/common/dto"
	dto2 "gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/dto"
	"net/http"
	"time"
)

type plantHireRequestService interface {
	Create(siteEngineerName string, constructionSiteName string, comment string, cost float64, plantInventoryEntryID int, startDate *time.Time, endDate *time.Time) (*domain.PlantHireRequest, error)
	GetAll() ([]*domain.PlantHireRequest, error)
}

// PlantHireRequestHandler is the API handler instances for plant hire requests
type PlantHireRequestHandler struct {
	plantHireRequestUC plantHireRequestService
}

// NewPlantHireRequestHandler returns new PlantHireRequestHandler API handler instance
func NewPlantHireRequestHandler(plantUC plantHireRequestService) *PlantHireRequestHandler {
	return &PlantHireRequestHandler{
		plantHireRequestUC: plantUC,
	}
}

// RegisterRoutes enables registering to existing mux router
func (h *PlantHireRequestHandler) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/phr").Subrouter()
	// POST /phr
	subRouter.HandleFunc("", h.createPlantHireRequest).Methods(http.MethodPost)
	// GET /phr
	subRouter.HandleFunc("", h.getPlantHireRequests).Methods(http.MethodGet)
}

// createPlantHireRequest handles creation of PHR
func (h *PlantHireRequestHandler) createPlantHireRequest(w http.ResponseWriter, r *http.Request) {

	var phr dto2.PlantHireRequest
	// defer func executes when calling return
	defer func() {
		if r.Body == nil {
			return
		}
		err := r.Body.Close()
		if err != nil {
			log.Errorf("Could not close request body, err %v", err)
		}
	}()
	err := json.NewDecoder(r.Body).Decode(&phr)
	if err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusBadRequest, "plant hire request format is invalid")
		return
	}

	if err = phr.Validate(); err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusInternalServerError, "error creating plant hire request")
		return
	}

	createdPhr, err := h.plantHireRequestUC.Create(phr.SiteEngineerName, phr.ConstructionSiteName, phr.Comment, phr.Cost, phr.PlantInventoryEntryID, phr.StartDate, phr.EndDate)
	if err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusInternalServerError, fmt.Sprintf("error creating plant hire request %v", phr))
		return
	}
	phrDto, err := dto2.NewPlantHireRequest(createdPhr)
	if err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusInternalServerError, "error creating plant hire request")
		return
	}
	// write success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&phrDto)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// getPlantHireRequests retrieves all PHR
func (h *PlantHireRequestHandler) getPlantHireRequests(w http.ResponseWriter, _ *http.Request) {
	phrs, err := h.plantHireRequestUC.GetAll()
	if err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve PHRs"))
		return
	}

	phrDtos, err := dto2.NewPlantHireRequests(phrs)
	if err != nil {
		log.Error(err.Error())
		writeErrorMessage(w, http.StatusInternalServerError, "error quering phrs")
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&phrDtos)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

// writeErrorMessage returns JSON error to API caller
func writeErrorMessage(w http.ResponseWriter, statusCode int, message string) {
	res := dto.NewError(message)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}
