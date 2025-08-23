package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/tickets", ListTickets)
		v1.GET("/tickets/:id", GetTicketByID)
		v1.POST("/tickets", CreateTicket)
		v1.DELETE("/tickets/:id", DeleteTicket)
	}
	r.Run("0.0.0.0:8080")

}
