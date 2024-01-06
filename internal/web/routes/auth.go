package routes

import (
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/middlewares"
)

func (r *Router) setAuth() {
	auth := r.router.Group("/auth")

	auth.Get(
		"/login",
		r.wrapper(middlewares.AlreadyLogin),
		r.wrapper(handlers.LoginPage),
	).Name("login")
	auth.Post(
		"/login",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.Login),
	)

	auth.Post(
		"/logout",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.Logout),
	).Name("logout")

	auth.Get(
		"/new",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(middlewares.AlreadyLogin),
		r.wrapper(handlers.RegisterPage),
	).Name("register")
	auth.Post(
		"/new",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.Register),
	)
}
