package initializers

import "github.com/pindarEng/stockSimulator/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
