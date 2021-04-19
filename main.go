package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/alexgtn/buildit/pkg/domain"
	commonHandlers "gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/common/handlers"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/handlers"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/infra/repository/db"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/service"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	cfg "gitlab.com/alexgtn/buildit/pkg/config"
	cfgLib "gitlab.com/alexgtn/buildit/pkg/config"

	// for SQL driver implementation
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	gormPG "gorm.io/driver/postgres"
)

const (
	configFilePath = "config"
	logLevel       = "debug"
)

func main() {
	// begin setup
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	// init config library
	loader := cfgLib.New()
	if err := loader.AddConfigPath("config"); err != nil {
		panic(err)
	}

	log.Infof("Config file path is %s, log level configured as %s ", configFilePath, logLevel)
	// load config file to struct
	var config cfg.Config
	if err := loader.Load(&config); err != nil {
		panic(err)
	}
	// print config
	log.Infof("%v", config)

	// Signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Info("Start buildit server api")

	// open Postgres connection
	DB, err := sql.Open("postgres", config.PostgreSQL.URI())
	if err != nil {
		log.Fatal(err)
	}
	// test connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("SQL Ping error=%v", err)
	}
	// initialize ORM using existing Postgres connection
	gormDB, err := gorm.Open(gormPG.New(gormPG.Config{
		Conn: DB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("ORM init error=%v", err)
	}
	// exec DB migrations
	err = gormDB.AutoMigrate(&domain.PlantHireRequest{})
	if err != nil {
		log.Fatalf("migrations error=%v", err)
	}
	// construct application
	plantHireRequestRepository := db.NewPlantHireRequestRepository(gormDB)
	plantHireRequestService := service.NewPlantHireRequestService(plantHireRequestRepository)
	plantHireRequestHandler := handlers.NewPlantHireRequestHandler(plantHireRequestService)

	router := mux.NewRouter()

	// register routes to router
	plantHireRequestHandler.RegisterRoutes(router)
	// ...

	// handler for not found routes
	router.NotFoundHandler = router.NewRoute().HandlerFunc(commonHandlers.NotFoundHandler).GetHandler()

	// setup http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTP.Service.Port),
		Handler: router,
	}

	// start server asynchronously in a goroutine
	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not start server")
		}
	}()

	log.Infof("All ready")

	// listen to system signals and shutdown when receiving such
	<-signals
	log.Info("Server shutdown")
	errSrvShutdown := srv.Shutdown(context.Background())
	if errSrvShutdown != nil {
		log.Fatalf("Could not shutdown server, err %v", err)
	}

}
