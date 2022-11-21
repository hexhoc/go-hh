package main

import (
	"hexhoc/go-hh/config"
	"hexhoc/go-hh/internal/app"
	"log"
	"os"
)

func main() {
	os.Getenv("HOME")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	app.Run(cfg)
}
