package Controllers

type controller interface {
	Mount(r *Router)
}
