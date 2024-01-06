package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmailRequired     = fiber.NewError(fiber.StatusBadRequest, "Email обязателен")
	ErrEmailIncorrect    = fiber.NewError(fiber.StatusBadRequest, "Введен некорректный email")
	ErrPasswordMissmatch = fiber.NewError(fiber.StatusBadRequest, "Пароли не совпадают")
	ErrPasswordRequired  = fiber.NewError(fiber.StatusBadRequest, "Пароль обязателен")
	ErrPasswordShort     = fiber.NewError(fiber.StatusBadRequest, "Пароль слишком короткий")
	ErrAuthorityRequired = fiber.NewError(fiber.StatusBadRequest, "Необходимо поле Округ/ЦОВУ")
	ErrUnitRequired      = fiber.NewError(fiber.StatusBadRequest, "Необходимо поле войсковая часть")
)

var userValidateMap = map[string]validatorFunc{
	"Email": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrEmailRequired
		case "email":
			return ErrEmailIncorrect
		}
		return nil
	},
	"Password": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrPasswordRequired
		case "min":
			return ErrPasswordShort
		}
		return nil
	},
	"PasswordConfirm": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrPasswordRequired
		case "eqfield":
			return ErrPasswordMissmatch
		}
		return nil
	},
	"Authority": func(e validator.FieldError) error {
		//nolint:revive //scalability
		switch e.Tag() {
		case required:
			return ErrAuthorityRequired
		}
		return nil
	},
	"Unit": func(e validator.FieldError) error {
		//nolint:revive //scalability
		switch e.Tag() {
		case required:
			return ErrUnitRequired
		}
		return nil
	},
}

func (v *Validator) SwitchUserValidate(user any) error {
	if err := v.validate.Struct(user); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint // This example from official doc
			for _, e := range err {
				validateFunc, ok := userValidateMap[e.Field()]
				if ok {
					if err := validateFunc(e); err != nil {
						return err
					}
				} else {
					return err
				}
				// if err := userValidateMap[e.Field()](e); err != nil {
				// 	return err
				// }
			}
			return err
		}
		return fiber.ErrInternalServerError
	}
	return nil
}
