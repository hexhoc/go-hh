package main

import (
	"fmt"
	"hexhoc/go-hh/config"
	"hexhoc/go-hh/internal/app"
	"log"
)

func main() {
	fmt.Println("START APP")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	app.Run(cfg)
}
