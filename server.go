package main

import "net/http"
import "fmt"

type Config struct {
	ListenAddr string
    StoreProducerFunc StoreProducerFunc
}

type Server struct {
	*Config
    topics map[string]Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
        Config: cfg,
        topics: make(map[string]Storer),
    }, nil
}

func (s *Server) Start() {
    http.ListenAndServe(s.ListenAddr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
}

func (s *Server) createTopic(name string) bool {
    _, ok := s.topics[name]
    if !ok {
        s.topics[name] = s.StoreProducerFunc()
        return true
    }
    return false
}
