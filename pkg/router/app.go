package router

import (
	"net/http"

	"api/pkg/middlewares"
	"api/pkg/service"
	"api/pkg/service/user"

	"github.com/go-chi/chi/v5"
)

// register app routes
func App(r chi.Router) {
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

	r.With(middlewares.Log).Route("/v1", func(r chi.Router) {
		r.Post("/login", service.Login)

		r.With(middlewares.Auth).Group(func(r chi.Router) {
			r.Get("/logout", service.Logout)

			{
				r.Get("/users", user.List)
				r.Post("/users", user.Create)
			}
		})
	})
}
