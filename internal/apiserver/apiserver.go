package apiserver

import (
	"io"
	"net/http"

	"github.com/AndreyGuznov/Avito_task2/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) StartServer() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting apiServer")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ApiServer) configureStore() error {
	storage := store.NewStore(s.config.Store)
	if err := storage.Open(); err != nil {
		return nil
	}

	s.store = storage

	return nil
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	// type request struct {
	// 	name string
	// }
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
