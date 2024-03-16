package app

import "net/http"

type App struct {
	httpServer *http.Server
}

func New(configPath string) (*App, error) {
	return &App{}, nil
}
