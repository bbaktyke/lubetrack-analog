package main

import (
	"bbaktyke/lubetrack-analog.git/internal/handlers"
	repo "bbaktyke/lubetrack-analog.git/internal/repository"
	"bbaktyke/lubetrack-analog.git/internal/service"
	"net/http"

	ftpserver "goftp.io/server/v2"
	"goftp.io/server/v2/driver/file"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	go func() {
		logrus.SetFormatter(new(logrus.JSONFormatter))
		if err := initConfig(); err != nil {
			logrus.Fatalf("error initializing configs:%s", err.Error())
		}
		db, err := connectToDatabase()
		if err != nil {
			logrus.Fatalf("failed to initialize db:%s", err.Error())
		}

		database, err := repo.New(db)
		if err != nil {
			logrus.Fatalf("failed to initialize database layer:%s", err.Error())
		}
		service, err := service.New(database)
		if err != nil {
			logrus.Fatalf("failed to initialize database layer:%s", err.Error())
		}
		server := handlers.New(service)
		if err != nil {
			logrus.Fatalf("failed to initialize server layer:%s", err.Error())
		}
		logrus.Printf("server was launched on port: %s", viper.GetString("port"))
		if err = http.ListenAndServe(viper.GetString("port"), server.Router); err != nil {
			logrus.Fatalf("failed to start server:%s", err.Error())
		}
	}()

	go func() {
		driver := &file.Driver{
			RootPath: "/srv/ftp",
		}

		// Creating the ftp server
		opts := &ftpserver.Options{
			Name:     "Lubetrack FTP Server",
			Driver:   driver,
			Port:     2121,
			Hostname: "localhost",
			Auth: &ftpserver.SimpleAuth{
				Name:     "admin",
				Password: "1234",
			},
			Perm:           ftpserver.NewSimplePerm("admin", "admin"),
			Logger:         new(ftpserver.DiscardLogger),
			WelcomeMessage: "Welcome to the Lubetrack FTP Server",
		}

		ftpServerInstance, err := ftpserver.NewServer(opts)
		if err != nil {
			logrus.Fatalf("Error creating ftp server: %v", err)
		}

		ftpServerInstance.RegisterNotifer(service.NotifierInstance)

		logrus.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
		if err := ftpServerInstance.ListenAndServe(); err != nil {
			logrus.Fatalf("Error starting ftp server: %v", err)
		}
	}()

	select {}
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
