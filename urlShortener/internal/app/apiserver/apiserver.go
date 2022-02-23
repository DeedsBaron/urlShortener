package apiserver

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"shortener/internal/app/config"
	"shortener/internal/app/store"
	"shortener/pkg/utils"
)

type APIServer struct {
	config  *config.Config
	logger  *logrus.Logger
	router  *mux.Router
	storage store.Storage
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: store.InitStorage(config),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Starting API server")
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
	s.router.HandleFunc("/encode/", s.CreateURL()).Methods("POST")
	s.router.HandleFunc("/{shortURL}", s.GetFullURL()).Methods("GET")
}

func (s *APIServer) GetFullURL() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := mux.Vars(r)["shortURL"]
		resp, err := s.storage.FindInStore(context.Background(), shortURL, s.config)
		if err != nil {
			utils.HttpErrorWithoutBackSlashN(w, err.Error(), http.StatusNotFound)
			s.logger.Info(err.Error())
			return
		}
		http.Redirect(w, r, resp, http.StatusSeeOther)
		s.logger.Debug("GET method SUCCESS")
	}
}

func (s *APIServer) CreateURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			LongUrl *string `json:"longUrl"` // pointer so we can test for field absence
		}{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&t)
		if err != nil {
			utils.HttpErrorWithoutBackSlashN(w, err.Error(), http.StatusBadRequest) //this func is needed for beautiful tests output
			s.logger.Error(err.Error())
			return
		}
		if t.LongUrl == nil {
			utils.HttpErrorWithoutBackSlashN(w, "Missing field 'longUrl' from JSON object", http.StatusBadRequest)
			s.logger.Error("Missing field 'longUrl' from JSON object")
			return
		}
		if d.More() {
			utils.HttpErrorWithoutBackSlashN(w, "Extraneous data after JSON object", http.StatusBadRequest)
			s.logger.Error("Extraneous data after JSON object")
			return
		}
		resp, err := s.storage.PostStore(context.Background(), *t.LongUrl, s.config)
		if err != nil {
			utils.HttpErrorWithoutBackSlashN(w, err.Error(), http.StatusBadRequest)
			s.logger.Error(err.Error())
			return
		}
		io.WriteString(w, resp)
		s.logger.Debug("POST method SUCCESS")
	}
}
