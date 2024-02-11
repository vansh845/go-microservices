package application

import (
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: loadRoutes(),
	}
}

func (a *App) Start() error {

	server := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}
	fmt.Println("server running at port 8000...🚀")
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server ,%d", err)
	}
	return err

}
