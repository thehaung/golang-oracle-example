package bootstrap

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/sijms/go-ora/v2"
	"github.com/thehaung/golang-oracle-example/config"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/database/oracle"
	"github.com/thehaung/golang-oracle-example/internal/httpserver"
	"github.com/thehaung/golang-oracle-example/repository"
	"github.com/thehaung/golang-oracle-example/transport/http/route"
	"github.com/thehaung/golang-oracle-example/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(conf *config.Config) {
	// Database
	oracleClient := oracle.NewOracleClient(conf)
	err := oracleClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	oracleConn, err := oracleClient.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	staffRepository := repository.NewStaffRepository(oracleConn)

	// UseCase
	staffUseCase := usecase.NewStaffUseCase(staffRepository)

	// HttpServer
	setupHttpServer(conf, staffUseCase)
}

func setupHttpServer(conf *config.Config, staffUseCase domain.StaffUseCase) {
	// HttpRouter
	router := chi.NewRouter()
	route.Setup(router, staffUseCase)
	httpServer := httpserver.NewServer(router, httpserver.Port(conf.HttpServer.Port))
	log.Println("HttpServer is serve at port", conf.HttpServer.Port)
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err := httpServer.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
