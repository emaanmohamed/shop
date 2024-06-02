package api

import (
	"database/sql"
	"github.com/emaanmohamed/shop/service/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	ListenAddr string
	DB         *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
		DB:         db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	sub := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.DB)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(sub)
	log.Println("JSON API server running on port: ", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, sub)
}
