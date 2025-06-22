package main

import (
	"fmt"
	"log"
	"os"

	"go-rails/framework/core"
	"go-rails/framework/generators"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorails",
	Short: "Go-Rails Framework CLI",
	Long:  `A command line tool for the Go-Rails framework with Rails-like functionality.`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the development server",
	Run: func(cmd *cobra.Command, args []string) {
		app := core.NewApplication()
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

var newCmd = &cobra.Command{
	Use:   "new [app_name]",
	Short: "Create a new Go-Rails application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		if err := generators.CreateNewApp(appName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created new Go-Rails application: %s\n", appName)
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate files for your application",
}

var generateControllerCmd = &cobra.Command{
	Use:   "controller [name]",
	Short: "Generate a new controller",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		if err := generators.GenerateController(controllerName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Generated controller: %s\n", controllerName)
	},
}

var generateModelCmd = &cobra.Command{
	Use:   "model [name] [fields...]",
	Short: "Generate a new model",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		fields := args[1:]
		if err := generators.GenerateModel(modelName, fields); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Generated model: %s\n", modelName)
	},
}

var generateMigrationCmd = &cobra.Command{
	Use:   "migration [name]",
	Short: "Generate a new migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrationName := args[0]
		if err := generators.GenerateMigration(migrationName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Generated migration: %s\n", migrationName)
	},
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands",
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		app := core.NewApplication()
		if err := app.DB.AutoMigrate(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database migrations completed successfully")
	},
}

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with sample data",
	Run: func(cmd *cobra.Command, args []string) {
		app := core.NewApplication()
		if err := generators.SeedDatabase(app.DB); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database seeded successfully")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(newCmd)

	generateCmd.AddCommand(generateControllerCmd)
	generateCmd.AddCommand(generateModelCmd)
	generateCmd.AddCommand(generateMigrationCmd)
	rootCmd.AddCommand(generateCmd)

	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbSeedCmd)
	rootCmd.AddCommand(dbCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
