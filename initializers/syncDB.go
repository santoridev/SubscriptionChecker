package initializers

import (
	"github.com/test/models"
)

func SyncDb() {
	DB.AutoMigrate(&models.Subs{})

}
