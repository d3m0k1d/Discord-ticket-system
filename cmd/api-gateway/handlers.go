package main

import (
	validation "Discord-ticket-system/internal/models/Validation"
	models "Discord-ticket-system/internal/models/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTickets(c *gin.Context) {
	var tickets []models.Ticket
	db.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}
func GetTicketByID(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
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
	id := c.Param("id")
	var ticket models.Ticket

	if err := db.Where("id = ?", id).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if err := db.Delete(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}

func PatchTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	if err := db.Where("id = ?", id).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
	}
	if ticket.Status == false {
		ticket.Status = true
		db.Save(&ticket)
		c.JSON(http.StatusOK, ticket)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Ticket already closest"})
	}
}
