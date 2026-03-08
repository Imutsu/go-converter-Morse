package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	HTTPServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	// создание http-роутер
	mux := http.NewServeMux() // mux - роутер

	// зарегистрировать handler-ы
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandlers)

	//создание самого http.Servee
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: logger,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10 *time.Second,
		IdleTimeout: 15 *time.Second,
	}

	// возвращение структуры Server с логгером
	return &Server{
		logger: logger,
		HTTPServer: httpServer,
	}
}