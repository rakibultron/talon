package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rakibultron/talon/db"
	"github.com/rakibultron/talon/routes"
	"github.com/spf13/cobra"
)

func main() {

	env := os.Getenv("APP_MODE")
	// Load environment variables
	envFile := fmt.Sprintf(".env.%s", env)
	envErr := godotenv.Load(envFile)
	if envErr != nil {

		fmt.Println(envErr)
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Create the root command
	var rootCmd = &cobra.Command{
		Use:   "talon",
		Short: "talon API Application",
		Long:  "talon API Application with support for embedded migrations and flexible commands.",
		Run: func(cmd *cobra.Command, args []string) {
			startServer(port)
		},
	}

	// Add migration commands
	rootCmd.AddCommand(db.MigrateCommand)

	// Execute the CLI
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func startServer(port string) {
	// Initialize Gin router
	r := gin.Default()
	routes.Routes(r)

	dbCon, dbconErr := db.DatabaseConnection()

	if dbconErr != nil {
		fmt.Println("Database connection failed!")
	} else {
		fmt.Println("Database connection success")
		fmt.Printf("%T", dbCon)
	}

	// Define basic ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	fmt.Printf("Starting server on port %s\n", port)
	r.Run(":" + port)
}
