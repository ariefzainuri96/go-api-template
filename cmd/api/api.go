package main

import (
	"log"
	"net/http"
	"time"

	middleware "github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	db   dbConfig
	addr string
}

type dbConfig struct {
	addr         string
	maxOpenCons  int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/v1/blog/", http.StripPrefix("/v1/blog", app.BlogRouter()))

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      stack(mux),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("Server has started on %s", app.config.addr)

	return srv.ListenAndServe()
}
