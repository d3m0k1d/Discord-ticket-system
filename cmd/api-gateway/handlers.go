package main

import (
	"Discord-ticket-system/internal/models/Validation"
	"github.com/gin-gonic/gin"
)

func ListTickets(c *gin.Context) {

}
func GetTicketByID(c *gin.Context) {

}
func CreateTicket(c *gin.Context) {
	var ticket = validation.TicketRequest{}
	c.ShouldBindJSON(&ticket)
	c.JSON(200, ticket)
}
func DeleteTicket(c *gin.Context) {

}
