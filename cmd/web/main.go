package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexandru-calin/tinycode/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	snippets      *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "web:password@/tinycode?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		snippets:      &models.SnippetModel{DB: db},
	}

	app.logger.Info("Starting the server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
