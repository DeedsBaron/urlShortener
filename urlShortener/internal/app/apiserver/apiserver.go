package apiserver

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"shortener/internal/app/store"
)

type APIServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage store.Storage
}

func New(config *Config) *APIServer {
	return &APIServer{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: store.InitStorage(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/encode/", s.createURL()).Methods("POST")
	s.router.HandleFunc("/{shortURL}", s.getFullURL()).Methods("GET")
}

func (s *APIServer) getFullURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := mux.Vars(r)["shortURL"]
		resp, err := s.storage.FindInStore(context.Background(), shortURL, s.config.Options.Schema, s.config.Options.Prefix)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, resp, http.StatusSeeOther)
		s.logger.Info("GET method SUCCESS")
	}
}

type request struct {
	url *string `json:"url"`
}

func (s *APIServer) createURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			LongUrl *string `json:"longurl"` // pointer so we can test for field absence
		}{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&t)
		if err != nil {
			// bad JSON or unrecognized json field
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.logger.Error(err.Error())
			return
		}
		if t.LongUrl == nil {
			http.Error(w, "Missing field 'longUrl' from JSON object", http.StatusBadRequest)
			s.logger.Error("Missing field 'longUrl' from JSON object")
			return
		}
		if d.More() {
			http.Error(w, "Extraneous data after JSON object", http.StatusBadRequest)
			s.logger.Error("Extraneous data after JSON object")
			return
		}
		resp, err := s.storage.PostStore(context.Background(), *t.LongUrl, s.config.Options.Schema, s.config.Options.Prefix)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.logger.Error(err.Error())
			return
		}
		io.WriteString(w, resp)
		s.logger.Info("POST method SUCCESS")
	}
}
