package handlers

import (
	"log"
	"net/http"

	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/storage"
	"github.com/gorilla/mux"
)

type HttpRouter struct {
	router *mux.Router
	db     storage.DbHandler
}

func NewHttpRouter(router *mux.Router, db storage.DbHandler) *HttpRouter {
	router.HandleFunc("/workplace", db.AddWorkplace).Methods(http.MethodPost)
	// router.HandleFunc("/plants", dbHandler.GetAllPlants).Methods(http.MethodGet)
	// router.HandleFunc("/plants", dbHandler.AddPlant).Methods(http.MethodPost)
	// router.HandleFunc("/plants/{id}", dbHandler.GetPlant).Methods(http.MethodGet)
	// router.HandleFunc("/plants/{id}", dbHandler.UpdatePlant).Methods(http.MethodPut)
	// router.HandleFunc("/plants/{id}", dbHandler.DeletePlant).Methods(http.MethodDelete)

	return &HttpRouter{router: router, db: db}
}

func (h *HttpRouter) Run(port string) {
	log.Fatal(http.ListenAndServe(port, h.router))
}
