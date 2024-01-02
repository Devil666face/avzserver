package handlers

import (
	"os"
	"path/filepath"

	"github.com/Devil666face/avzserver/internal/web/view"
	"github.com/Devil666face/avzserver/pkg/file"
)

func Bases(h *Handler) error {
	base, err := os.Getwd()
	if err != nil {
		return h.c.Next()
	}
	abs, err := filepath.Abs(filepath.Join(base, h.c.Path()))
	if err != nil {
		return h.c.Next()
	}
	if stat, err := os.Stat(abs); err != nil || !stat.IsDir() {
		return h.c.Next()
	}
	files, err := file.DirContent(abs)
	if err != nil {
		return h.c.Next()
	}
	return h.Render(view.BasesList, view.Map{
		view.DirContentKey: files,
	})
}
