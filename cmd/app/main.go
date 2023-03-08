package main

import (
	"log"

	"github.com/hezzl_task5/config"
	"github.com/hezzl_task5/internal/app"
)

func main() {
	// get config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// run app
	app.Run(cfg)
}
