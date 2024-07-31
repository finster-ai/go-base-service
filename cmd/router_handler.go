package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

// muxHandler wraps a mux.Router to meet go-micro's Handler interface requirements
type muxHandler struct {
	name   string
	router *mux.Router
}

func (m *muxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m *muxHandler) Handler() interface{} {
	return m
}

func (m *muxHandler) Endpoints() []*registry.Endpoint {
	return []*registry.Endpoint{}
}

func (m *muxHandler) Options() server.HandlerOptions {
	return server.HandlerOptions{}
}

func (m *muxHandler) Name() string {
	return m.name
}

func newMuxHandler(name string, router *mux.Router) *muxHandler {
	return &muxHandler{
		name:   name,
		router: router,
	}
}
