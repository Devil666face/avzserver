package middlewares

import (
	"strings"

	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func AllowHost(h *handlers.Handler) error {
	if host, ok := h.Ctx().GetReqHeaders()[view.Host]; ok {
		if strings.Contains(host[0], h.Config().AllowHost) {
			return h.Ctx().Next()
		}
	}
	return fiber.ErrBadRequest
}
