package initializers

import "jwt/models"

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}