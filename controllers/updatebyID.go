package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test/initializers"
	"github.com/test/models"
	"gorm.io/gorm"
)

func UpdateByID(c *gin.Context) {
	ID := c.Param("id")

	var input models.Subs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if input.ServiceName == "" || input.Price <= 0 || input.UserID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service name, price or user_id"})
		return
	}

	var sub models.Subs
	result := initializers.DB.First(&sub, "id = ?", ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		} else {
			log.Printf("DB error fetching subscription %s: %v", ID, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	sub.ServiceName = input.ServiceName
	sub.Price = input.Price
	//sub.UserID = input.UserID если нужно можно включить, но думаю менять user_id не надо. :)

	result = initializers.DB.Save(&sub)
	if result.Error != nil {
		log.Printf("DB error updating subscription %s: %v", ID, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subscription"})
		return
	}

	log.Printf("Subscription %s updated successfully", ID)
	c.JSON(http.StatusOK, sub)

}
