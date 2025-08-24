package main

import (
	"Discord-ticket-system/internal/models/db"
	"fmt"
	"log"
	_ "net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Ticket{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func main() {

	initDB()

	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/tickets", ListTickets)
		v1.GET("/tickets/:id", GetTicketByID)
		v1.POST("/tickets", CreateTicket)
		v1.DELETE("/tickets/:id", DeleteTicket)
		v1.PATCH("/tickets/:id", PatchTicket)
	}
	r.Run("0.0.0.0:8080")

}
