package main

import (
	validation "Discord-ticket-system/internal/models/Validation"
	models "Discord-ticket-system/internal/models/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTickets(c *gin.Context) {

}
func GetTicketByID(c *gin.Context) {

}
func CreateTicket(c *gin.Context) {
	var req validation.TicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket := models.Ticket{
		Title:       req.Title,
		Description: req.Description,
	}
	if err := db.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}
func DeleteTicket(c *gin.Context) {

}
