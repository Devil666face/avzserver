package view

import (
	"fmt"
	"log/slog"
	"path"
	"strings"

	"github.com/Devil666face/avzserver/internal/models"
	"github.com/gofiber/fiber/v2"
)

const (
	UserID       = "user_id"
	Csrf         = "csrf"
	Htmx         = "htmx"
	HxRequest    = "Hx-Request"
	HxCurrentURL = "Hx-Current-Url"
	HXRedirect   = "HX-Redirect"
	Host         = "Host"
	// HXRefresh    = "HX-Refresh"
)

const (
	AuthKey           = "authenticated"
	UserKey           = "User"
	UsersKey          = "Users"
	AllertMessageKey  = "AlertMessage"
	SuccessMessageKey = "SuccessMessage"
	DirContentKey     = "Content"
)

type View struct {
	*fiber.Ctx
}

func New(c *fiber.Ctx) *View {
	return &View{c}
}

func (c View) CsrfToken() string {
	if token, ok := c.Locals(Csrf).(string); ok {
		return token
	}
	return ""
}

func (c View) URL(name string) string {
	return c.getRouteURL(name, fiber.Map{})
}

func (c View) URLto(name, key string, val any) string {
	return c.getRouteURL(name, fiber.Map{
		key: val,
	})
}

func (c View) IsURL(name string) bool {
	return c.Ctx.OriginalURL() == c.URL(name)
}

func (c View) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		slog.Error(fmt.Sprintf("url %s not found", name))
	}
	return url
}

func (c View) ActivateURL(u models.User) string {
	return c.BaseURL() + c.getRouteURL("user_activate", fiber.Map{
		"u":   u.Email,
		"otp": u.OneTimeCode,
	})
}

func (c View) IsHtmx() bool {
	if htmx, ok := c.Locals(Htmx).(bool); ok {
		return htmx
	}
	return false
}

func (c View) SetClientRedirect(redirectURL string) error {
	c.Set(HXRedirect, redirectURL)
	return c.SendStatus(fiber.StatusFound)
}

func (c View) PreviousPage() string {
	s := strings.Split(path.Clean(c.Path()), "/")
	if len(s) == 2 {
		return "/"
	}
	return strings.Join(s[0:len(s)-1], "/")
}

func (c View) CurrentUser() models.User {
	if u, ok := c.Locals(UserKey).(models.User); ok {
		return u
	}
	return models.User{}
}

func (c View) IsCurrentURL(url string) bool {
	return c.Path() == url
}

// func (c ViewCtx) SetClientRefresh() {
// 	c.Set(HXRefresh, "true")
// }

// func (c ViewCtx) IsHtmxCurrentURL() bool {
// 	if url, ok := c.GetReqHeaders()[HxCurrentURL]; ok {
// 		return url[0] == c.BaseURL()+c.OriginalURL()
// 	}
// 	return false
// }
