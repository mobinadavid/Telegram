package pgx

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"os"
	"sync"
	"telegram/src/database/config"
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
		configs := config.GetInstance() // Retrieve configurations
		db, err = pgx.Connect(context.Background(),
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
				configs.Get("DB_USERNAME"),
				configs.Get("DB_PASSWORD"),
				configs.Get("DB_HOST"),
				configs.Get("DB_PORT"),
				configs.Get("DB_DATABASE"),
			),
		)
	})

	return err
}

func Close() error {
	if db == nil {
		return nil
	}

	return db.Close(context.Background())
}

func RunFileQuery(file string) (pgconn.CommandTag, error) {
	// Read the SQL file
	sqlQuery, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Execute the SQL query
	return db.Exec(context.Background(), string(sqlQuery))
}
