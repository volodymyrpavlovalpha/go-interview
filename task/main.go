package main

import (
	"context"
	"log"

	"task/inmem"
	"task/rates"
	"task/web"
)

func main() {
	// TODO: Initialize accounts handler in a correct way and run a program

	err := handler.ServeHTTP(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
