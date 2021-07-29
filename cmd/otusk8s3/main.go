package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"

	"github.com/adalekin/otusk8s3/internal/api/http/server"
	"github.com/adalekin/otusk8s3/internal/appenv"
	"github.com/adalekin/otusk8s3/internal/repository/reporegistry"
	"github.com/caarlos0/env"
)

func main() {
	log.Info("Starting the service...")

	// Parse environment variables
	config := appenv.Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Create a database object
	dbConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName)
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a repository registry object
	repoRegistry := reporegistry.NewPostgreSQL(db)

	// Initialize an application environment
	appEnv := appenv.AppEnv{Config: config, RepoRegistry: repoRegistry, Render: render.New()}

	// Run a HTTP server
	srv := server.MakeServer(appEnv)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Info("The service is ready to listen and serve.")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Info("Got SIGINT...")
	case syscall.SIGTERM:
		log.Info("Got SIGTERM...")
	}

	log.Info("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Info("Done")
}
