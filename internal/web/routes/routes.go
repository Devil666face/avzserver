package routes

import (
	"github.com/Devil666face/avzserver/assets"
	"github.com/Devil666face/avzserver/internal/config"
	"github.com/Devil666face/avzserver/internal/mail"
	"github.com/Devil666face/avzserver/internal/store/database"
	"github.com/Devil666face/avzserver/internal/store/session"
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/middlewares"
	"github.com/Devil666face/avzserver/internal/web/validators"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	router      fiber.Router
	config      *config.Config
	database    *database.Database
	store       *session.Store
	validator   *validators.Validator
	mail        *mail.Smtp
	middlewares []func(*handlers.Handler) error
}

func New(
	_router fiber.Router,
	_config *config.Config,
	_database *database.Database,
	_store *session.Store,
	_validator *validators.Validator,
	_mail *mail.Smtp,
) *Router {
	r := Router{
		router:    _router,
		config:    _config,
		database:  _database,
		store:     _store,
		validator: _validator,
		mail:      _mail,
		middlewares: []func(*handlers.Handler) error{
			middlewares.Recover,
			middlewares.Logger,
			// middlewares.Compress,
			middlewares.Limiter,
			middlewares.AllowHost,
			middlewares.SecureHeaders,
			middlewares.EncryptCookie,
			middlewares.Csrf,
			middlewares.Htmx,
		},
	}
	r.setMiddlewares()
	r.setAuth()
	r.setUser()
	r.setBases()
	r.setIndex()
	return &r
}

func (r *Router) wrapper(handler func(*handlers.Handler) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return handler(handlers.New(c, r.database, r.config, r.store, r.validator, r.mail))
	}
}

func (r *Router) setMiddlewares() {
	r.router.Use(r.wrapper(middlewares.Compress))
	r.router.Use(assets.DirStatic, r.wrapper(middlewares.Static))
	for _, middleware := range r.middlewares {
		r.router.Use(r.wrapper(middleware))
	}
}
