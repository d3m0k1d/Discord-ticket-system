package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/tickets", ListTickets)
		v1.GET("/tickets/:id", GetTicketByID)
		v1.POST("/tickets", CreateTicket)
		v1.DELETE("/tickets/:id", DeleteTicket)
	}
	r.Run(":8080")

}
