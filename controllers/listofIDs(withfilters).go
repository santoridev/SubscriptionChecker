package controllers

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test/initializers"
	"github.com/test/models"
)

func ListOfIDs(c *gin.Context) {
	ID := c.Query("user_id")
	serviceName := c.Query("service_name")
	userID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "It's not an uuid",
		})
		return
	}
	if serviceName == "" || len(serviceName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Service name is missing"})
		return
	}
	var subs []models.Subs

	result := initializers.DB.Find(&subs, "user_id = ? AND service_name = ?", userID, serviceName)

	if result.Error != nil {
		if result.Error != nil {
			log.Printf("Database error: %v", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find subscriptions"})
			return
		}

	}
	log.Printf("Found %d subscriptions for user %s and service %s", len(subs), userID, serviceName)
	c.JSON(http.StatusOK, subs)
}
