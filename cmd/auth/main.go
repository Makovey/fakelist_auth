package main

import (
	syslog "log"

	"github.com/Makovey/fakelist_auth/internal/app"
	"github.com/Makovey/fakelist_utils/pkg/logger/slog"
)

func main() {
	log := slog.NewLogger()

	appl := app.NewApp(log)

	if err := appl.Run(); err != nil {
		syslog.Fatal(err)
	}
}
