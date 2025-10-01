package controllers

import (
	"log"

	"github.com/aloneen/assign2-ex2-adv/initializers"
	"github.com/aloneen/assign2-ex2-adv/models"
)

func CreateUserWithProfile(user models.User) {
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		log.Println("Insert failed:", result.Error)
	} else {
		log.Println("Inserted user with profile:", user.ID)
	}
}

func GetUsersWithProfiles() {
	var users []models.User
	err := initializers.DB.Preload("Profile").Find(&users).Error
	if err != nil {
		log.Println("Query failed:", err)
		return
	}

	for _, u := range users {
		log.Printf("User: %s (Age %d) -> Profile: %s\n", u.Name, u.Age, u.Profile.Bio)
	}
}

func UpdateUserProfile(userID uint, newBio string) {
	err := initializers.DB.Model(&models.Profile{}).
		Where("user_id = ?", userID).
		Update("bio", newBio).Error

	if err != nil {
		log.Println("Update failed:", err)
	} else {
		log.Println("Profile updated")
	}
}

func DeleteUser(userID uint) {
	err := initializers.DB.Delete(&models.User{}, userID).Error
	if err != nil {
		log.Println("Delete failed:", err)
	} else {
		log.Println("User (and profile) deleted")
	}
}
