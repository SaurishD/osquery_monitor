package main

import (
	"log"

	"github.com/SaurishD/osquery_monitor/internal/api"
	"github.com/SaurishD/osquery_monitor/internal/db"
	"github.com/SaurishD/osquery_monitor/internal/osquery"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Init() // connect to DB

	osInfo, err := osquery.GetOSInfo()
	if err != nil {
		log.Fatalf("Failed to get OS info: %v", err)
	}

	osqVer, err := osquery.GetOsqueryVersion()
	if err != nil {
		log.Fatalf("Failed to get osquery version: %v", err)
	}

	apps, err := osquery.GetInstalledApps()
	if err != nil {
		log.Fatalf("Failed to get installed apps: %v", err)
	}

	if err := db.InsertLatestSnapshot(database, osInfo, osqVer, apps); err != nil {
		log.Fatalf("Failed to insert snapshot: %v", err)
	}

	router := gin.Default()
	router.GET("/latest_data", api.GetLatestData)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
