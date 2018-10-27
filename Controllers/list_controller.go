package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
	"github.com/DigitalAnswer/PlaygroundTodoList/models"
	"github.com/DigitalAnswer/PlaygroundTodoList/services"
	"github.com/rs/zerolog/log"
)

// ListController struct
type ListController struct {
	listService *services.ListService
}

// NewListController constructor
func NewListController(service *services.ListService) (*ListController, error) {
	return &ListController{
		listService: service,
	}, nil
}

// Mount routes
func (c ListController) Mount(r *Router) {
	r.AddRouteFunc("/list", c.Add).Methods(http.MethodPost)
}

// Add func
func (c *ListController) Add(w http.ResponseWriter, r *http.Request) {
	req := &models.List{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Err(err).Msg("Unable to decode")
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
		return
	}

	r.Body.Close()
	list, err := c.listService.Create(req)
	if err != nil {
		log.Error().Err(err).Msg("Cannot add a list")
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	jsonResponse := map[string]interface{}{
		"list": list,
	}

	JSON(w, http.StatusOK, jsonResponse)
}
