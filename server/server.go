package server

import (
    "fmt"
    "net/http"
)

const (
    Port int    = 9871
    Ip   string = "127.0.0.1"
)

type Server struct {
    Addr   string
    Id     int
    Master bool
}

func Create(id int, addr string, master bool) Server {
    return Server{addr, id, master}
}

func (s *Server) Start() {
    http.ListenAndServe(s.Addr, nil)
    // http.Handle("/user/get", handler)

    fmt.Printf("server starting on %s", s.Addr)
}
