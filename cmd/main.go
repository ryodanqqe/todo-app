package main

import (
	"log"

	"github.com/ryodanqqe/todo-app"
	"github.com/ryodanqqe/todo-app/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run server due to error: %s", err.Error())
	}
}
