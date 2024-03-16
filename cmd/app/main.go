package main

import (
	"context"
	"flag"
	"log"
	"time"

	"diplom-backend/internal/app"
)

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config-path", "./config/config.yaml", "path to configuration file")
	flag.Parse()

	appl, err := app.New(configPath)
	if err != nil {
		log.Fatalf("Initializing app: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = appl.Run(ctx); err != nil {
		log.Fatalf("Running an application: %v\n", err)
	}
}
