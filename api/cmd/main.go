package main

import (
	"context"
	"fmt"
	"go-temp/go-sample/api/internal/app"
	"log"
	"time"
)

func main() {
	local, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("failed to load location: %v\n", err)
	}
	time.Local = local

	ctx := context.Background()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %v\n", err)
	}

	fmt.Println(app)

	defer func() {
		if err := app.Close(); err != nil {
			log.Fatalf("failed to close app: %v\n", err)
		}
	}()

	if err := app.Migrate(); err != nil {
		log.Printf("failed to migrate: %v\n", err)
	}

	if err := app.Start(); err != nil {
		log.Printf("failed to start app: %v\n", err)
	}

	log.Println("migrated")
}
