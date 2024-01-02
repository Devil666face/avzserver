package routes

import (
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/middlewares"
)

func (r *Router) setIndex() {
	index := r.router.Group("/")
	index.Use(r.wrapper(middlewares.Auth))

	index.Get(
		"",
		r.wrapper(handlers.Index),
	).Name("index")
}
