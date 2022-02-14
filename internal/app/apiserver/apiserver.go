package apiserver

import (
	"encoding/json"
	"fmt"
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
	s.logger.Info("starting api server")
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
	s.router.HandleFunc("/", s.handleHello())
	s.router.HandleFunc("/post", s.createURL()).Methods("POST")
	//s.router.HandleFunc("/get", s.getFullURL(shortURL string)).Methods("GET")
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

type request struct {
	url *string `json:"url"`
}

func (s *APIServer) createURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			Url *string `json:"url"` // pointer so we can test for field absence
		}{}
		d := json.NewDecoder(r.Body)
		//d.DisallowUnknownFields()
		err := d.Decode(&t)
		if err != nil {
			// bad JSON or unrecognized json field
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if t.Url == nil {
			http.Error(w, "missing field 'test' from JSON object", http.StatusBadRequest)
			return
		}
		if d.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
			return
		}
		fmt.Println("long url=", *t.Url)
		resp := s.storage.PostStore(*t.Url, s.config.Options.Schema, s.config.Options.Prefix)
		//makeShortUrl()
		io.WriteString(w, s.config.Options.Schema+s.config.Options.Prefix+resp)
	}
}
