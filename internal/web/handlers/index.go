package handlers

import (
	"github.com/Devil666face/avzserver/internal/web/view"
)

func Index(h *Handler) error {
	return h.Render(view.Index, view.Map{})
}
