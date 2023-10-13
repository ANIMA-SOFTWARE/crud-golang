package api

import (
	"main/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	router     mux.Router
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		router:     *mux.NewRouter().StrictSlash(false),
	}
}

func (s *Server) Start() error {

	//.Methods("GET")
	s.router.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	s.router.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	s.router.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))

	s.router.HandleFunc("/", s.handleBase)
	s.router.HandleFunc("/templates/{template}", s.handleTemplates)
	s.router.HandleFunc("/scripts/{script}", s.handleScripts)
	s.router.HandleFunc("/stylesheets/{stylesheet}", s.handleStylesheets)

	s.router.HandleFunc("/data/{table}", s.handleDataGetFirst).Methods("GET")
	s.router.HandleFunc("/data/{table}/id={id}", s.handleDataGetByID).Methods("GET")
	s.router.HandleFunc("/data/{table}/page={PageNumber}", s.handleDataGetPage).Methods("GET")
	s.router.HandleFunc("/data/{table}/search", s.handleDataGetBySearch).Methods("POST")
	s.router.HandleFunc("/data/{table}/id={id}/edit", s.handleDataEditByID).Methods("GET")
	s.router.HandleFunc("/data/{table}/id={id}", s.handleDataDeleteByID).Methods("DELETE")
	s.router.HandleFunc("/data/{table}/id={id}", s.handleDataAppendByID).Methods("PUT")
	s.router.HandleFunc("/data/{table}/new", s.handleDataCreateByID).Methods("POST")

	return http.ListenAndServe(s.listenAddr, &s.router)
}
