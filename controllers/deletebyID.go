package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test/initializers"
	"github.com/test/models"
)

func DeleteByID(c *gin.Context) {
	idParam := c.Param("id")

	result := initializers.DB.Delete(&models.Subs{}, "id = ?", idParam)
	if result.Error != nil {
		log.Printf("DB error deleting subscription %s: %v", idParam, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	log.Printf("Subscription %s deleted successfully", idParam)
	c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted"})
}
