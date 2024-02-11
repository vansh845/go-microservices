package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	return &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
}

func (a *App) Start() error {

	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}
	err := a.rdb.Ping(context.Background()).Err()
	fmt.Println("connected to redis...")
	if err != nil {
		return fmt.Errorf("failed to connect to redis server...%s", err.Error())
	}
	fmt.Println("server running at port 8000...🚀")
	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server ,%d", err)
	}
	return err

}
