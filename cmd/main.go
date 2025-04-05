package main

import (
	"log"

	"github.com/SaurishD/osquery_monitor/internal/api"
	"github.com/SaurishD/osquery_monitor/internal/db"
	"github.com/SaurishD/osquery_monitor/internal/osquery"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	database := db.Init() // connect to DB

	c := cron.New()

	// Runs every 30 seconds
	_, err := c.AddFunc("@every 30s", func() {
		osInfo, err := osquery.GetOSInfo()
		if err != nil {
			log.Printf("Error getting OS info: %v", err)
			return
		}

		osqVer, err := osquery.GetOsqueryVersion()
		if err != nil {
			log.Printf("Error getting osquery version: %v", err)
			return
		}

		apps, err := osquery.GetInstalledApps()
		if err != nil {
			log.Printf("Error getting installed apps: %v", err)
			return
		}

		if err := db.InsertLatestSnapshot(database, osInfo, osqVer, apps); err != nil {
			log.Printf("Error inserting snapshot: %v", err)
		} else {
			log.Println("Snapshot saved successfully.")
		}
	})
	if err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}

	c.Start()

	router := gin.Default()
	router.GET("/latest_data", api.GetLatestData)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
