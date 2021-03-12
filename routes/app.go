package routes

import (
	"net/http"

	"github.com/shenghui0779/demo/controllers"
	"github.com/shenghui0779/demo/middlewares"

	"github.com/go-chi/chi/v5"
)

// RegisterApp register app routes
func RegisterApp(r chi.Router) {
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

	// r.Method(http.MethodGet, "/metrics", promhttp.Handler())

	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.Logger)

		r.Route("/books", func(r chi.Router) {
			r.Get("/info/{id}", controllers.BookInfo)
		})
	})
}
