package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Devil666face/avzserver/internal/web/view"
	"github.com/Devil666face/avzserver/pkg/utils"
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
	files, err := utils.DirContent(abs)
	if err != nil {
		return h.c.Next()
	}
	fmt.Println(files)
	return h.Render(view.Index, view.Map{})
}
