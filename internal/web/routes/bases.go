package routes

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/avzserver/assets"
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/middlewares"
	"github.com/Devil666face/avzserver/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) setBases() {
	bases := r.router.Group("/bases")
	bases.Use(r.wrapper(middlewares.Auth))

	path, err := utils.SetDir(assets.DirBases)
	if err != nil {
		slog.Error(fmt.Sprintf("Media directory not create or found: %s", err))
		//nolint:revive //If dir for bases not created or not have - close app
		os.Exit(1)
	}
	bases.Use(r.wrapper(handlers.Bases))

	bases.Static("", path, fiber.Static{
		Download:  true,
		ByteRange: true,
	})
}
