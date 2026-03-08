package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// создание логгера
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// создание сервера с помощью функции NewServer
	srv := server.NewServer(logger) // srv - указатель на нашу структуру

	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
