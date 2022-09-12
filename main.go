package main

import (
	"routekey/database"
	"routekey/models"
	"routekey/routes"
	"routekey/config"
	
	"github.com/gin-gonic/gin"

	"time"
	"log"
	"strconv"
)

var startTime time.Time = time.Now()

func main() {
	cfg := config.ReadConfig()
	
	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	
	_ = database.Initialize()

	database.AuthDB.AutoMigrate(&models.User{})
	database.LinksDB.AutoMigrate(&models.Link{})
	database.LinksDB.AutoMigrate(&models.Domain{})
	database.LinksDB.AutoMigrate(&models.Tracker{})

	routes.StartTime = startTime
	server := routes.Setup()
	routes.BootTime = time.Since(startTime)
	log.Printf("Routekey V2 started on %s\n", cfg.Port)
	log.Printf("Loaded %s words.\n", strconv.Itoa(len(cfg.Words)))
	log.Fatal(server.Run(cfg.Port))
}