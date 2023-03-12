package main

import (
	"log"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/app"
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
