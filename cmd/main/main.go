package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/avzserver/internal/web"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	wa := web.New()
	if err := wa.Listen(); err != nil {
		slog.Error(fmt.Sprintf("start programm: %s", err))
		os.Exit(1)
	}
}
