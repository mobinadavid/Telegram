package bootstrap

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"telegram/src/database/cockroach"
	"telegram/src/database/pgx"
	"time"
)

func Init() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v. ", err)
	}
	log.Printf("ENV Service: Env variable initial successfully.\n")

	// Initialize pgx:Database
	err = pgx.Init()
	if err != nil {
		log.Fatalf("Database:pgx Service: Failed to Initialize: %v.", err)
	}
	log.Printf("Database:pgx Service: Database  initial successfully. \n")

	//
	err = cockroach.Init()
	if err != nil {
		log.Fatalf("Database:cockroach Service: Failed to Initialize: %v.", err)
	}
	log.Printf("Database:cockroach Service: Database  initial successfully. \n")

	// app started ...
	time.Sleep(50 * time.Millisecond)
	log.Printf("Application is now running.Press CTRL-C to exit.\n")
	<-sc

	// Shutting down application
	log.Printf("Application shutting down....   \n")

	// Close Database:pgx
	err = pgx.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v. \n", err)
	}
	log.Printf("Databasde Service: database close sucessfully. \n")

	// wait for every thing down
	time.Sleep(1 * time.Second)
}
