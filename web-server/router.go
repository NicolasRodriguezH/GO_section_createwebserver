package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exist := r.FindHandler(req.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, req)
}
