package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test/initializers"
	"github.com/test/models"
)

func GetSubSummary(c *gin.Context) {
	ID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start_date")
	end := c.Query("end_date")

	userID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	startDate, err := time.Parse("01-2006", start)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}
	endDate, err := time.Parse("01-2006", end)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	var total int

	err = initializers.DB.
		Model(&models.Subs{}).
		Where("user_id = ? AND service_name = ? AND start_date BETWEEN ? AND ?", userID, serviceName, startDate, endDate).
		Select("SUM(price)").Scan(&total).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to calculate total"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
