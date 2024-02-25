package main

import (
	"bbaktyke/lubetrack-analog.git/internal/handlers"
	repo "bbaktyke/lubetrack-analog.git/internal/repository"
	"bbaktyke/lubetrack-analog.git/internal/service"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}

func run() error {
	database, err := repo.New()
	if err != nil {
		return err
	}
	service, err := service.New(database)
	if err != nil {
		return err
	}
	server := handlers.New(service)
	if err != nil {
		return err
	}
	log.Printf("server was launched on port: %d", 8080)
	return http.ListenAndServe(":8080", server.Router)
}
