package routes

import (
	"CursoGolang/myedcoments/controllers"

	"github.com/gorilla/mux"
)

// SetLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
