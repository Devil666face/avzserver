package handlers

import (
	"github.com/Devil666face/avzserver/internal/models"
	"github.com/Devil666face/avzserver/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInSession    = fiber.ErrInternalServerError
	ErrDisabledUser = fiber.NewError(fiber.StatusBadRequest, "Пользователь не активен")
)

func LoginPage(h *Handler) error {
	return h.Render(view.Login, view.Map{})
}

func Login(h *Handler) error {
	var (
		u   = &models.User{}
		in  = &models.User{}
		err error
	)
	if err := h.Ctx().BodyParser(in); err != nil {
		return err
	}
	u.Email = in.Email
	if err := u.LoginValidate(h.Database(), h.Validator(), in.Password); err != nil {
		return h.Render(view.Login, view.Map{
			view.MessageKey: err.Error(),
		})
	}
	if !u.Active {
		return h.Render(view.Login, view.Map{
			view.MessageKey: ErrDisabledUser.Error(),
		})
	}
	if err := h.SetInSession(view.AuthKey, true); err != nil {
		return ErrInSession
	}
	if err := h.SetInSession(view.UserID, u.ID); err != nil {
		return ErrInSession
	}
	if u.SessionKey, err = h.SessionID(); err != nil {
		return ErrInSession
	}
	if err := u.Update(h.Database()); err != nil {
		return ErrInSession
	}
	return h.View().ClientRedirect(h.View().URL("index"))
}

func Logout(h *Handler) error {
	if err := h.DestroySession(); err != nil {
		return ErrInSession
	}
	return h.View().ClientRedirect(h.View().URL("login"))
}

func RegisterPage(h *Handler) error {
	return h.Render(view.Register, view.Map{})
}

func Register(h *Handler) error {
	u := models.User{}
	if err := h.Ctx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	return nil
}
