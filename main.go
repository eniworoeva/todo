package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/db"
	"todolist/handlers"
	"todolist/repository"
	"todolist/routes"
	"todolist/services"
)

func main() {
	db.InitDb()

	// initialize repositories
	userRepo := &repository.UserRepo{}
	todoRepo := &repository.TodoRepo{}

	// initialize service
	userService := &services.UserService{Repo: userRepo}
	todoService := &services.TodoService{Repo: todoRepo}

	// initialize handlers
	userHandler := &handlers.UserHandler{userService}
	todoHandler := &handlers.TodoHandler{todoService}

	//routes
	r := routes.SetupRouter(userHandler, todoHandler)

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}
