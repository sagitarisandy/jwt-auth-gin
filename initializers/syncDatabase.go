package initializers

import (
	"github.com/sagitarisandy/jwt-auth-gin/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
