package router

import (
	"net/http"

	"api/app/pkg/controller"
	"api/app/pkg/middleware"
	"api/app/web"
	"api/lib"
	lib_middleware "api/lib/middleware"

	"github.com/go-chi/chi/v5"
)

// register app routes
func App(r chi.Router) {
	lib.FileServer(r, "/", http.FS(web.Asserts()))

	// 浏览器访问会主动发送 /favicon.ico 请求
	// r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "favicon.ico")
	// })

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// prometheus metrics
	// r.Method(http.MethodGet, "/metrics", promhttp.Handler())

	r.With(lib_middleware.Log).Group(func(r chi.Router) {
		// v1
		r.Route("/v1", func(r chi.Router) {
			v1(r)
		})
	})
}

func v1(r chi.Router) {
	r.Post("/login", controller.Login)

	r.With(middleware.Auth).Group(func(r chi.Router) {
		r.Get("/logout", controller.Logout)

		{
			r.Post("/user/list", controller.UserList)
			r.Post("/user/info", controller.UserInfo)
			r.Post("/user/create", controller.UserCreate)
		}
	})
}
