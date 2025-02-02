package cockroach

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"sync"
)

var (
	once sync.Once
	db   *pgx.Conn
)

func Init() error {
	return connect()
}

func GetInstance() *pgx.Conn {
	return db
}

func connect() error {
	var err error

	once.Do(func() {
		// Hardcoded values for connection
		username := "root"         // Replace with your actual username
		password := "yourpassword" // Replace with your actual password
		host := "localhost"        // Replace with your host if different
		port := "26257"            // Replace with your port if different
		database := "defaultdb"    // Replace with your database name if different

		// Connect to CockroachDB using the JDBC-like URL
		// JDBC URL: jdbc:postgresql://localhost:26257/defaultdb
		db, err = pgx.Connect(context.Background(),
			fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
				username,
				password,
				host,
				port,
				database,
			),
		)

		// Check connection
		if err != nil {
			fmt.Println("Unable to connect to CockroachDB:", err)
			return
		}

		fmt.Println("Connected to CockroachDB")
	})

	return err
}

func Close() error {
	if db == nil {
		return nil
	}

	return db.Close(context.Background())
}
