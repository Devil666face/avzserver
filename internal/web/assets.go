package web

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/avzserver/assets"
	"github.com/Devil666face/avzserver/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

const (
	isBrowse = false
)

type Media struct {
	path    string
	handler fiber.Static
}

// func NewStatic() func(*fiber.Ctx) error {
// 	return filesystem.New(filesystem.Config{
// 		Root:       http.FS(assets.StaticFS),
// 		PathPrefix: assets.DirStatic,
// 		MaxAge:     86400,
// 	})
// }

func MustMedia() *Media {
	path, err := utils.SetDir(assets.DirMedia)
	if err != nil {
		slog.Error(fmt.Sprintf("Media directory not create of found: %s", err))
		//nolint:revive //If dir for media not created - close app
		os.Exit(1)
	}
	return &Media{
		path: path,
		handler: fiber.Static{
			Compress:  true,
			ByteRange: true,
			Browse:    isBrowse,
		},
	}
}
