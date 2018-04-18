package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router struct
type Router struct {
	*mux.Router
	middleware alice.Chain
}

// NewRouter constructor
func NewRouter() *Router {
	return &Router{
		Router: mux.NewRouter(),
	}
}

// Use middleware
func (r *Router) Use(middleware ...alice.Constructor) {
	r.middleware = r.middleware.Append(middleware...)
}

// AddController to router
func (r *Router) AddController(c controller) {
	c.Mount(r)
}

// AddRoute to router
func (r *Router) AddRoute(path string, handler http.Handler) *mux.Route {
	return r.Handle(path, r.middleware.Then(handler))
}

// AddRouteFunc to router
func (r *Router) AddRouteFunc(path string, handler http.HandlerFunc) *mux.Route {
	return r.Handle(path, r.middleware.ThenFunc(handler))
}
