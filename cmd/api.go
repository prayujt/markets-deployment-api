package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"markets-api/internal/handlers"
)

func main() {
	router := mux.NewRouter()
	router.Use(recoverMiddleware)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	handler := handlers.NewClientEnvironment()
	defer handler.QueueClient.Close()
	log := handler.Log
	router.HandleFunc("/markets/deploy", handler.QueueDeploymentRequest).Methods("POST")
	router.HandleFunc("/markets/{marketID}", handler.GetDeploymentStatus).Methods("GET")

	port := os.Getenv("API_PORT")
	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("starting server", "port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Error("server listen failed", "err", err)
	}
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
