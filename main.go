package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	m "github.com/junichi-tanaka/go-chi-zap/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(m.NewStructuredLogger(logger))
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oops")
	})
	http.ListenAndServe(":3333", r)
}
