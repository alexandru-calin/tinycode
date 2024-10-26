package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	app.logger.Info("Starting the server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
