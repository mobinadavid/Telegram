package database

import (
	"github.com/spf13/cobra"
	"telegram/cmd/database/migrate"
	"telegram/cmd/database/seed"
)

// Database Commands for interacting with apps
var DatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Commands for interacting with database.",
}

func init() {
	DatabaseCmd.AddCommand(
		migrate.Migrate,
		migrate.MigrateCockroach,
		seed.Seed,
	)
}
