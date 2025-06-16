package main

import (
	syslog "log"

	"github.com/Makovey/fakelist_auth/internal/app"
	"github.com/Makovey/fakelist_utils/pkg/config/env"
	"github.com/Makovey/fakelist_utils/pkg/logger/slog"
)

func main() {
	fn := "main"

	log := slog.NewLogger()

	cfg, err := env.NewAuthConfig(log)
	if err != nil {
		syslog.Fatalf("[%s]: could not parse config: %s", fn, err)
	}

	appl := app.NewApp(log, cfg)

	if err = appl.Run(); err != nil {
		syslog.Fatalf("[%s]: app run failed cause: %s", fn, err)
	}
}
