package initializers

import (
	"time"

	"github.com/google/uuid"
	"github.com/test/models"
	"gorm.io/gorm"
)

func InsertFixedSubscription(db *gorm.DB) error {
	userID, _ := uuid.Parse("60601fee-2bf1-4721-ae6f-7636e79a0cba")
	startDate, _ := time.Parse("01-2006", "07-2025")

	sub := models.Subs{
		UserID:      userID,
		ServiceName: "Yanddddex Plus",
		Price:       4100,
		StartDate:   models.MonthYear(startDate),
		EndDate:     nil,
	}

	return db.Create(&sub).Error
}
