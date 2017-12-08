package main

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
	mux map[string]func(http.ResponseWriter, *http.Request)
}

func NewHandler() *Handler {
	Handler := &Handler{}
	Handler.mux = make(map[string]func(http.ResponseWriter, *http.Request))
	return Handler
}

func (Handler *Handler) Bind(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handler.mux[pattern] = handler
}

func (Handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zero := time.Now()
	defer func() {
		duration := time.Since(zero)
		log.Printf("%s [%.2f] %s \n", r.Method, duration.Seconds(), r.URL)
	}()
	if h, ok := Handler.mux[r.URL.String()]; ok {
		h(w, r)
	} else {
		w.WriteHeader(404)
	}
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
}

func main() {
	Handler := NewHandler()
	Handler.Bind("/get", HandlerGet)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: Handler,
	}
	server.ListenAndServe()
}
