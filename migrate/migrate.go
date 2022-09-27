package main

import (
	"fmt"
	"log"

	"github.com/vtmhieu/golang-gorm-postgres/initializers"
	"github.com/vtmhieu/golang-gorm-postgres/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}

//we loaded the environment variables and created a connection pool to the Postgres database in the init() function
//Then, we evoked the AutoMigrate() function provided by GORM to create the database migration and push the changes to the database.
