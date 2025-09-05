package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test/initializers"
	"github.com/test/models"
)

func GetListOfIDs(c *gin.Context) {
	var subs []models.Subs

	result := initializers.DB.Find(&subs)
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}

	c.JSON(http.StatusOK, subs)
}
