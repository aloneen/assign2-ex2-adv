package main

import (
	"log"

	"github.com/aloneen/assign2-ex2-adv/initializers"
	"github.com/aloneen/assign2-ex2-adv/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		panic(err)
	}
	log.Println("Database migrated")
}
