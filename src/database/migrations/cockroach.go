package migrations

import (
	"context"
	"fmt"
	"os"
	"telegram/src/database/cockroach"
	"telegram/src/utils"
)

func MigrateUpCockroach() error {
	// Read all SQL files in the "migrations" directory
	files, err := utils.ReadSQLFiles("src/database/migrations/cockroach/up")
	if err != nil {
		return err
	}

	// Execute each migration file
	for i := 0; i < len(files); i++ {
		sqlQuery, err := os.ReadFile(files[i])
		if err != nil {
			return err
		}
		_, err = cockroach.GetInstance().Exec(context.Background(), string(sqlQuery))
		if err != nil {
			fmt.Println(i)
			return err
		}
	}
	return nil
}

func MigrateDownCockroach() error {
	// Read all SQL files in the "migrations" directory
	files, err := utils.ReadSQLFiles("src/database/migrations/down")
	if err != nil {
		return err
	}

	// Execute each migration file
	for i := len(files) - 1; i >= 0; i-- {
		sqlQuery, err := os.ReadFile(files[i])
		if err != nil {
			return err
		}
		_, err = cockroach.GetInstance().Exec(context.Background(), string(sqlQuery))
		if err != nil {
			return err
		}
	}
	return nil
}
