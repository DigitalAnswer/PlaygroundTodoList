package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
	"github.com/DigitalAnswer/PlaygroundTodoList/models"
	"github.com/DigitalAnswer/PlaygroundTodoList/services"
	"github.com/gorilla/mux"
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
	r.AddRouteFunc("/list/all", c.GetAllList).Methods(http.MethodGet)
	r.AddRouteFunc("/list/delete", c.Delete).Methods(http.MethodDelete)
	r.AddRouteFunc("/list/{id}", c.Get).Methods(http.MethodGet)
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

	userID := r.Context().Value(helpers.KeyPrincipalID).(int64)

	list, err := c.listService.Create(userID, req)
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

// GetAllList func
func (c *ListController) GetAllList(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(helpers.KeyPrincipalID).(int64)

	allList, err := c.listService.GetAllList(userID)
	if err != nil {
		log.Error().Err(err).Msg("Cannot get all list")
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	jsonResponse := map[string]interface{}{
		"lists": allList,
	}

	JSON(w, http.StatusOK, jsonResponse)
}

// Delete list func
func (c *ListController) Delete(w http.ResponseWriter, r *http.Request) {

	type deleteList struct {
		ID int64 `json:"id"`
	}

	req := &deleteList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Err(err).Msg("Unable to decode")
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
		return
	}
	r.Body.Close()

	userID := r.Context().Value(helpers.KeyPrincipalID).(int64)
	err := c.listService.Delete(userID, req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Cannot delete list")
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	jsonResponse := map[string]interface{}{}

	JSON(w, http.StatusOK, jsonResponse)
}

// Get func
func (c *ListController) Get(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(helpers.KeyPrincipalID).(int64)

	listIDString := mux.Vars(r)["id"]
	listID, err := strconv.ParseInt(listIDString, 10, 64)
	if err != nil {
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	list := &models.List{}
	list, err = c.listService.Get(userID, listID)
	if err != nil {
		log.Error().Err(err).Msg("Cannot get list by id")
		helpers.FailureFromError(w, http.StatusForbidden, err)
		return
	}

	jsonResponse := map[string]interface{}{
		"list": list,
	}

	JSON(w, http.StatusOK, jsonResponse)
}
