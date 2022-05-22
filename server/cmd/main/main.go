package main

import (
	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/handlers"
	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/storage"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	dbHandler := storage.NewDBHandler(storage.NewDB())

	httpRouter := handlers.NewHttpRouter(router, dbHandler)

	httpRouter.Run(":8080")

}
