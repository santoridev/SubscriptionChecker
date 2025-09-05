package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test/initializers"
	"github.com/test/models"
)

func GetByID(c *gin.Context) {
	ID := c.Param("id")

	var subs []models.Subs
	result := initializers.DB.Where("id = ?", ID).Find(&subs)
	if result.Error != nil {
		log.Printf("DB error while fetching subscriptions  %s: %v", ID, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if len(subs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No subscriptions found "})
		return
	}

	c.JSON(http.StatusOK, subs)
}
