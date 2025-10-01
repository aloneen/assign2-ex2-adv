package main

import (
	"github.com/aloneen/assign2-ex2-adv/controllers"
	"github.com/aloneen/assign2-ex2-adv/initializers"
	"github.com/aloneen/assign2-ex2-adv/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {

	// Create a new user
	//newUser := models.User{
	//	Name: "Dias",
	//	Age:  19,
	//	Profile: models.Profile{
	//		Bio:               "Go developer",
	//		ProfilePictureURL: "http://example.com/dias.png",
	//	},
	//}

	user2 := models.User{
		Name: "Dosbol",
		Age:  21,
		Profile: models.Profile{
			Bio:               "Java developer",
			ProfilePictureURL: "http://example.com/dosbol.png",
		},
	}
	controllers.CreateUserWithProfile(user2)

	// Query
	controllers.GetUsersWithProfiles()

	// Update
	controllers.UpdateUserProfile(1, "Updated Bio")

	// Delete
	controllers.DeleteUser(1)
}
