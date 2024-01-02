package handlers

import (
	"os"
	"path/filepath"

	"github.com/Devil666face/avzserver/internal/web/view"
	"github.com/Devil666face/avzserver/pkg/file"
)

func URLToFilepath(url string) (string, error) {
	base, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(filepath.Join(base, url))
	if err != nil {
		return "", err
	}
	return abs, nil
}

func Bases(h *Handler) error {
	path, err := URLToFilepath(h.c.Path())
	if err != nil {
		return h.c.Next()
	}
	if stat, err := os.Stat(path); err != nil || !stat.IsDir() {
		return h.c.Next()
	}
	files, err := file.DirContent(path)
	if err != nil {
		return h.c.Next()
	}
	return h.Render(view.BasesList, view.Map{
		view.DirContentKey: files,
	})
}
