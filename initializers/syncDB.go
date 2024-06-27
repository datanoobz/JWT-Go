package initializers

import "jwt-go/models"

func SyncDB() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
