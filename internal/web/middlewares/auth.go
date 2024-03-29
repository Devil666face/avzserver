package middlewares

import (
	"github.com/Devil666face/avzserver/internal/models"
	"github.com/Devil666face/avzserver/internal/web/handlers"
	"github.com/Devil666face/avzserver/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func Auth(h *handlers.Handler) error {
	var (
		u   = models.User{}
		uID any
		err error
		ok  bool
	)
	if auth, err := h.GetFromSession(view.AuthKey); auth == nil || err != nil {
		return h.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if uID, err = h.GetFromSession(view.UserID); uID == nil || err != nil {
		return h.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if u.ID, ok = uID.(uint); !ok {
		return h.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if err := u.Get(h.Database()); err != nil {
		return h.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	h.Ctx().Locals(view.UserKey, u)
	return h.Ctx().Next()
}

func AlreadyLogin(h *handlers.Handler) error {
	auth, err := h.GetFromSession(view.AuthKey)
	if auth, ok := auth.(bool); auth && ok && err == nil {
		return h.Ctx().RedirectToRoute("index", nil)
	}
	return h.Ctx().Next()
}
