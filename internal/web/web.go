package web

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/avzserver/assets"
	"github.com/Devil666face/avzserver/internal/config"
	"github.com/Devil666face/avzserver/internal/mail"
	"github.com/Devil666face/avzserver/internal/models"
	"github.com/Devil666face/avzserver/internal/store/database"
	"github.com/Devil666face/avzserver/internal/store/session"
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/routes"
	"github.com/Devil666face/avzserver/internal/web/validators"

	"github.com/gofiber/fiber/v2"
)

type Web struct {
	fiber     *fiber.App
	media     *Media
	config    *config.Config
	database  *database.Database
	router    *routes.Router
	store     *session.Store
	validator *validators.Validator
	mail      *mail.Mail
	tables    []any
}

func New() *Web {
	a := &Web{
		fiber: fiber.New(
			fiber.Config{
				ErrorHandler:          handlers.DefaultErrorHandler,
				DisableStartupMessage: true,
			},
		),
		media:     MustMedia(),
		config:    config.Must(),
		validator: validators.New(),
		tables: []any{
			&models.User{},
		},
	}
	a.setStores()
	a.setMail()
	a.setStatic()
	a.setRoutes()
	return a
}

func (a *Web) setStores() {
	a.database = database.Must(a.config, a.tables)
	a.store = session.New(a.config, a.database)
}

func (a *Web) setMail() {
	a.mail = mail.New(a.config)
}

func (a *Web) setStatic() {
	a.fiber.Static(assets.DirMedia, a.media.path, a.media.handler)
}

func (a *Web) setRoutes() {
	a.router = routes.New(a.fiber, a.config, a.database, a.store, a.validator, a.mail)
}

func (a *Web) Listen() error {
	if a.config.UseTLS {
		return a.listenTLS()
	}
	return a.listenNoTLS()
}

func (a *Web) listenTLS() error {
	go a.mustRedirectServer()
	return a.fiber.ListenTLS(a.config.ConnectHTTPS, a.config.TLSCrt, a.config.TLSKey)
}

func (a *Web) listenNoTLS() error {
	return a.fiber.Listen(a.config.ConnectHTTP)
}

func (a *Web) mustRedirectServer() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect(a.config.HTTPSRedirect)
	})
	if err := app.Listen(a.config.ConnectHTTP); err != nil {
		slog.Error(fmt.Sprintf("Start redirect server: %s", err))
		//nolint:revive //If connection for redirect server already busy - close app
		os.Exit(1)
	}
}
