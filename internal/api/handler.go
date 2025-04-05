package api

import (
	"net/http"

	"github.com/SaurishD/osquery_monitor/internal/db"
	"github.com/gin-gonic/gin"
)

// internal/api/handler.go

func GetLatestData(c *gin.Context) {
	osInfo, apps, err := db.FetchLatestSnapshot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"os_info":      osInfo,
		"applications": apps,
	})
}
