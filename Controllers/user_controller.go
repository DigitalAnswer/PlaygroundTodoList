package Controllers

import (
	"net/http"
)

// UserController struct
type UserController struct {
}

// NewUserController constructor
func NewUserController() (*UserController, error) {
	return &UserController{}, nil
}

// Mount routes
func (c UserController) Mount(r *Router) {
	r.AddRouteFunc("/login", c.Login).Methods(http.MethodGet)
	r.AddRouteFunc("/login", c.Login).Methods(http.MethodPost)
	r.AddRouteFunc("/logout", c.Logout).Methods(http.MethodPost)
}

// Login func
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
}

// Logout func
func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
}
