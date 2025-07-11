package routes

import (
	"todolist/handlers"
	"todolist/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// //sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// //authenticated routes
	// protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	protected.HandleFunc("/todos", todoHandler.CreateToDo).Methods("POST")
	protected.HandleFunc("/todos", todoHandler.GetToDos).Methods("GET")
	// protected.HandleFunc("/todos/{id}", handlers.UpdateToDo).Methods("PUT")
	// protected.HandleFunc("/todos/{id}", handlers.DeleteToDo).Methods("DELETE")

	return r

}
