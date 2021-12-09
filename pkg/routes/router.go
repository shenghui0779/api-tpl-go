package routes

import (
	"net/http"
	"tplgo/pkg/middlewares"
	"tplgo/pkg/service"

	"github.com/go-chi/chi/v5"
)

// Register register routes
func Register(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("☺ welcome to golang app"))
	})

	// 浏览器访问会主动发送 /favicon.ico 请求
	// r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	return
	// })

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// prometheus metrics
	// r.Method(http.MethodGet, "/metrics", promhttp.Handler())

	r.With(middlewares.Logger).Route("/v1/", func(r chi.Router) {
		user(r)
	})
}

func user(r chi.Router) {
	user := service.NewUser()

	r.Get("/users/{id}", user.Info)
}
