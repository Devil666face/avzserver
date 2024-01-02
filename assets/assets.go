package assets

import (
	"embed"
)

const (
	DirMedia  = "media"
	DirStatic = "static"
	DirBases  = "bases"
)

//go:embed static/*
var StaticFS embed.FS
