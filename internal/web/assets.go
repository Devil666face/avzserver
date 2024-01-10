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
