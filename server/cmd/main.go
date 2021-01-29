package main

import (
	"context"
	"log"

	"module12/task05/server/internal/system/database/psql"
)

func main() {
	// Клиент для psql
	_, err := psql.New(context.Background(), "postgres://module12_task05:module12_task05@localhost:5432/module12_task05?sslmode=disable&connect_timeout=5")
	if err != nil {
		log.Fatal(err)
	}


	log.Println("Hello from task 05 sever!")
}
