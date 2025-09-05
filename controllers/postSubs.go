// AddSubs добавляет новую подписку
// @Summary Создание подписки
// @Description Добавляет новую подписку пользователя
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body models.Subs true "Данные подписки"
// @Success 201 {object} models.Subs
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test/initializers"
	"github.com/test/models"
)

func AddSubs(c *gin.Context) {
	var req models.Subs

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.ServiceName == "" || req.Price <= 0 || req.UserID == uuid.Nil {
		log.Printf("%s, %d, %s", req.ServiceName, req.Price, req.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input fields"})
		return
	}

	StoredReq := models.Subs{
		UserID:      req.UserID,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	result := initializers.DB.Create(&StoredReq)
	if result.Error != nil {
		log.Println("DB error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store data"})
		return
	}

	c.JSON(http.StatusCreated, StoredReq)
}
