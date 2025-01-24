package main

import (
	"log"
	"net/http"
	"simpleBank/src/api/handler"
	"simpleBank/src/api/router"
	"simpleBank/src/repository"
	"simpleBank/src/usecase"
)

func main() {
	initServer()
}

func initServer() {
	//setup all router
	payaHandler := setupPayaHandler()
	router := router.SetupPayaRouter(payaHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupPayaHandler() *handler.PayaHandler {
	payaRepo := repository.NewInMemoryPayaRepository()
	payaUseCase := usecase.NewShebaUseCase(payaRepo)

	return handler.NewPayaHandler(payaUseCase)
}
