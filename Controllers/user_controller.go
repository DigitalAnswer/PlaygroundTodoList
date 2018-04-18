package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/services"

	"github.com/rs/zerolog/log"

	"github.com/DigitalAnswer/PlaygroundTodoList/models"
)

// UserController struct
type UserController struct {
	userService *services.UserService
}

// NewUserController constructor
func NewUserController(service *services.UserService) (*UserController, error) {
	return &UserController{
		userService: service,
	}, nil
}

// Mount routes
func (c UserController) Mount(r *Router) {
	r.AddRouteFunc("/signin", c.Signin).Methods(http.MethodGet)
	r.AddRouteFunc("/signin", c.Signin).Methods(http.MethodPost)
	r.AddRouteFunc("/signout", c.Signout).Methods(http.MethodPost)
	r.AddRouteFunc("/signup", c.Signup).Methods(http.MethodPost)
}

// Signin func
func (c *UserController) Signin(w http.ResponseWriter, r *http.Request) {

}

// Signout func
func (c *UserController) Signout(w http.ResponseWriter, r *http.Request) {
}

// Signup func
func (c *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Err(err).Msg("Unable to decode")

		FailureFromError(w, http.StatusInternalServerError, err)

		return
	}

	defer r.Body.Close()
	user, err := c.userService.Create(req)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create new user")
		FailureFromError(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, user)
}
