package route

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thehaung/golang-oracle-example/domain"
	"log"
	"net/http"
	"time"
)

func Setup(router *chi.Mux, staffUC domain.StaffUseCase) {
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(15 * time.Second))

	// Prometheus' metrics register
	router.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})

	// Register this router for health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal("Hello World!!!!")
		if err != nil {
			log.Fatal(fmt.Errorf("NewRouter - json.Marshal: %w", err))
		}
		if _, err = w.Write(res); err != nil {
			log.Println(err.Error())
		}
	})

	// Global prefix
	router.Mount("/api/v1", router)
	{
		newStaffRoute(router, staffUC)
	}
}
