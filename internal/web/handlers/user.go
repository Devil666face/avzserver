package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Devil666face/avzserver/internal/models"
	"github.com/Devil666face/avzserver/internal/web/validators"
	"github.com/Devil666face/avzserver/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func UserListPage(h *Handler) error {
	// if h.View().IsHtmx() {
	// 	return h.Render(view.UserContent, view.Map{
	// 		view.UsersKey: models.GetAllUsers(h.Database()),
	// 	})
	// }
	return h.Render(view.UserList, view.Map{
		view.UsersKey: models.GetAllUsers(h.Database()),
	})
}

func UserEditForm(h *Handler) error {
	u := models.User{}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	u.Password = ""
	return h.Render(view.UserEdit, view.Map{
		view.UserKey: u,
	})
}

func UserCreateForm(h *Handler) error {
	return h.Render(view.UserCreate, view.Map{})
}

func UserCreate(h *Handler) error {
	u := models.User{}
	if err := h.Ctx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(h.Validator()); err != nil {
		return h.Render(view.UserCreate, view.Map{
			view.UserKey:          u,
			view.AllertMessageKey: err.Error(),
		})
	}
	if err := u.Create(h.Database()); err != nil {
		return h.Render(view.UserCreate, view.Map{
			view.UserKey:          u,
			view.AllertMessageKey: err.Error(),
		})
	}
	return h.Render(view.UserCreate, view.Map{
		view.SuccessMessageKey: fmt.Sprintf("User %s - created", u.Email),
	})
}

func UserEdit(h *Handler) error {
	var (
		u  = models.User{}
		in = models.User{}
	)
	if err := h.Ctx().BodyParser(&in); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	in.ID = uint(id)
	u.ID = in.ID
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := in.Validate(h.Validator()); err != nil {
		if errors.Is(err, validators.ErrPasswordRequired) {
			in.Password, in.PasswordConfirm = u.Password, u.Password
		} else {
			return h.Render(view.UserEdit, view.Map{
				view.UserKey:          u,
				view.AllertMessageKey: err.Error(),
			})
		}
	}
	u.Email, u.Password, u.Authority, u.Unit, u.Admin, u.Active = in.Email, in.Password, in.Authority, in.Unit, in.Admin, in.Active
	if err := u.Update(h.Database()); err != nil {
		return h.Render(view.UserEdit, view.Map{
			view.UserKey:          u,
			view.AllertMessageKey: err.Error(),
		})
	}
	if err := h.DestroySessionByID(u.SessionKey); err != nil {
		return h.Render(view.UserEdit, view.Map{
			view.UserKey:          u,
			view.AllertMessageKey: err.Error(),
		})
	}
	return h.Render(view.UserEdit, view.Map{
		view.UserKey:           u,
		view.SuccessMessageKey: "Пользователь обновлен",
	})
}

func UserDelete(h *Handler) error {
	u := models.User{}
	if err := h.Ctx().BodyParser(&u); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := u.Delete(h.Database()); err != nil {
		return err
	}
	if err := h.DestroySessionByID(u.SessionKey); err != nil {
		return ErrInSession
	}
	return h.Render(view.UserContent, view.Map{
		view.UsersKey: models.GetAllUsers(h.Database()),
	})
}

func UserActivate(h *Handler) error {
	var (
		u = models.User{}
	)
	email, otp := h.Ctx().Params("u"), h.Ctx().Params("otp")
	if email == "" || otp == "" {
		return fiber.ErrNotFound
	}
	if err := u.GetByUsername(h.Database(), email); err != nil {
		return fiber.ErrNotFound
	}
	if u.Active {
		return fiber.ErrNotFound
	}
	if u.OneTimeCode != otp {
		return fiber.ErrNotFound
	}
	u.Active, u.OneTimeCode = true, ""
	if err := u.Update(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	return h.Render(view.Login, view.Map{
		view.UserKey:           u,
		view.SuccessMessageKey: fmt.Sprintf("Пользователь %s - активирован", u.Email),
	})
}
