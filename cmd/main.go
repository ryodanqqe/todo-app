package main

import (
	"log"

	"github.com/ryodanqqe/todo-app"
	"github.com/ryodanqqe/todo-app/pkg/handler"
	"github.com/ryodanqqe/todo-app/pkg/repository"
	"github.com/ryodanqqe/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to initialize configs due to error: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run server due to error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
