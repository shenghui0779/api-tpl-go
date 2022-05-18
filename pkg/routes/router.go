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

	r.Route("/v1/", func(r chi.Router) {
		r.With(middlewares.Logger()).Group(func(r chi.Router) {
			{
				s := service.NewAuth()

				r.Post("/login", s.Login)
				r.With(middlewares.Auth).Get("/logout", s.Logout)
			}

			r.With(middlewares.Auth).Group(func(r chi.Router) {
				{
					s := service.NewUser()

					r.Post("/users", s.Create)
					r.Get("/users", s.List)
				}
			})
		})
	})
}
