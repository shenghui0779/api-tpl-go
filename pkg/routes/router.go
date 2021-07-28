package routes

import (
	"net/http"

	"tplgo/internal/middlewares"
	"tplgo/pkg/routes/v1"

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

	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewares.Logger)

		r.Route("/users", func(r chi.Router) {
			r.Get("/info/{id}", v1.UserInfo)
		})
	})
}
