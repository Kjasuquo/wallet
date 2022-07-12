package config

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type config struct {
	Port     int
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewConfig() *config {
	return &config{}
}

func (cfg *config) Server() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: cfg.routes(),
	}

	cfg.InfoLog.Printf("Running on port %d", cfg.Port)
	return srv.ListenAndServe()

}
