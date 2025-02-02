package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"log"
	"os"
	"sync"
	"telegram/src/database/cockroach"
	"time"
)

//func main() {
//	err := cmd.Execute()
//	if err != nil {
//		panic(err)
//	}
//}

const (
	// Number of concurrent operations to simulate workload
	numOperations = 10
)

func main() {
	// Step 1: Initialize the connection to CockroachDB
	err := cockroach.Init()
	if err != nil {
		log.Fatalf("Error initializing CockroachDB connection: %v", err)
	}

	// Step 2: Start a WaitGroup to wait for all operations to complete
	var wg sync.WaitGroup

	// Step 3: Generate workload by running multiple operations concurrently
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go simulateDatabaseOperation(i, &wg)
	}

	// Step 4: Wait for all operations to finish
	wg.Wait()

	// Step 5: Close the database connection when done
	err = cockroach.Close()
	if err != nil {
		log.Fatalf("Error closing CockroachDB connection: %v", err)
	}

	fmt.Println("Workload simulation completed.")
}

// Simulate a database operation (e.g., insert, select, update)
func simulateDatabaseOperation(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when done

	// Example: Run a query (e.g., insert data or perform a transaction)
	// You can modify this to perform different types of queries

	// Simple SQL command to insert data (simulated workload)
	sqlQuery := fmt.Sprintf("INSERT INTO users (id, name, email) VALUES ('%d', 'User %d', 'user%d@example.com');", id, id, id)

	// Run the query
	_, err := runFileQuery(sqlQuery) // Directly running SQL instead of reading from a file
	if err != nil {
		log.Printf("Error executing query %d: %v", id, err)
		return
	}

	// Sleep to simulate workload time
	time.Sleep(100 * time.Millisecond) // Simulate some delay between operations
}

// RunFileQuery reads an SQL file and executes the queries inside it.
func runFileQuery(file string) (pgconn.CommandTag, error) {
	// Step 1: Read the SQL file
	sqlQuery, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", file, err)
	}

	// Step 2: Execute the SQL query using the CockroachDB connection
	return cockroach.GetInstance().Exec(context.Background(), string(sqlQuery))
}
