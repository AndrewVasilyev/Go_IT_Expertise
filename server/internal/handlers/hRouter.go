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
	router.HandleFunc("/workplace", db.GetWorkplace).Methods(http.MethodGet)
	router.HandleFunc("/workplace", db.UpdateWorkplace).Methods(http.MethodPut)
	router.HandleFunc("/workplace", db.DeleteWorkplace).Methods(http.MethodDelete)

	return &HttpRouter{router: router, db: db}
}

func (h *HttpRouter) Run(port string) {
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(port, h.router))
}
