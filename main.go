package main

import (
	"fmt"
	"log"
)

const (
	awsSecretAccessKey = "wJa1rAf5dnFEMi/K7MDENG/BPRfiC1BYw28qMnQ+EV5"
)

func main() {
	fmt.Println("Demo Secrets Application")
	fmt.Println("========================")

	if err := connectToAWS(); err != nil {
		log.Fatalf("Fatal: failed to connect to AWS: %v", err)
	}

	if err := connectToMongoDB(); err != nil {
		log.Fatalf("Fatal: failed to connect to MongoDB: %v", err)
	}

	fmt.Println("\nApplication started successfully!")
}

func connectToAWS() error {
	fmt.Printf("Connecting to AWS with key %s...\n", awsSecretAccessKey[0:4]+"****")
	return nil
}

func connectToMongoDB() error {
	fmt.Printf("Connecting to MongoDB at %s...\n", "prod-cluster.example.com")
	connStr := GetMongoDBConnectionString()
	_ = connStr
	return nil
}
