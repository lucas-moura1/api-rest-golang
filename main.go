package main

import (
	"context"
	"fmt"

	"log"
	"net/http"

	"user-api/config"
	"user-api/implementation"
	"user-api/infrastructure/router"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

    userController := implementation.GetImplementation(ctx, cfg.Database.Name, cfg.Database.Domain)
    r := router.NewRouter(userController)

	log.Printf("Server listen in port: %s", cfg.HTTP.Port)
    http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTP.Port), r)
}
