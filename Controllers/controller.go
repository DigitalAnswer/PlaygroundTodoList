package controllers

type controller interface {
	Mount(r *Router)
}
