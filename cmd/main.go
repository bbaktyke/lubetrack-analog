package main

import (
	"bbaktyke/lubetrack-analog.git/internal/handlers"
	repo "bbaktyke/lubetrack-analog.git/internal/repository"
	"bbaktyke/lubetrack-analog.git/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}

func run() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs:%s", err.Error())
		return err
	}
	db, err := connectToDatabase()
	if err != nil {
		logrus.Fatalf("failed to initialize db:%s", err.Error())
		return err
	}

	database, err := repo.New(db)
	if err != nil {
		logrus.Errorf("failed to initialize database layer:%s", err.Error())
		return err
	}
	service, err := service.New(database)
	if err != nil {
		logrus.Errorf("failed to initialize database layer:%s", err.Error())
		return err
	}
	server := handlers.New(service)
	if err != nil {
		logrus.Errorf("failed to initialize server layer:%s", err.Error())
		return err
	}
	logrus.Printf("server was launched on port: %s", viper.GetString("port"))
	return http.ListenAndServe(viper.GetString("port"), server.Router)
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func connectToDatabase() (*sqlx.DB, error) {
	return repo.NewPostgresDB(repo.Confiq{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
}
