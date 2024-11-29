package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// MigrateCommand is the root command for migrations
var MigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration commands",
	Long:  "Commands for managing database migrations, including up, down, reset, and status.",
}

// init function to add subcommands
func init() {
	MigrateCommand.AddCommand(migrateUpCommand)
	MigrateCommand.AddCommand(migrateDownCommand)
	MigrateCommand.AddCommand(migrateStatusCommand)
}

// migrateUpCommand applies all migrations
var migrateUpCommand = &cobra.Command{
	Use:   "up",
	Short: "Apply all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations("up")
	},
}

// migrateDownCommand rolls back the last migration
var migrateDownCommand = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations("down")
	},
}

// migrateStatusCommand checks migration status
var migrateStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "Check migration status",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations("status")
	},
}

func runMigrations(action string) {

	// Configure database connection
	databaseURL := os.Getenv("DB_URI")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Configure goose to use embedded migrations
	goose.SetBaseFS(embedMigrations)

	// Execute goose command with context
	if err := goose.RunContext(ctx, action, db, "migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	fmt.Printf("Migrations %s completed successfully\n", action)
}
