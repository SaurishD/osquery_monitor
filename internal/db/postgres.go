package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// Init initializes the database connection
func Init() {
	err := godotenv.Load(".env") // or just godotenv.Load() if it's in root
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connStr := getDatabaseURL()

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&OSInfo{}, &Application{})
	if err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}
}

func InsertLatestSnapshot(osInfo map[string]string, osqueryVer map[string]string, apps []map[string]string) error {
	// Start transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Insert OSInfo
	newOS := OSInfo{
		Platform:   osInfo["platform"],
		Version:    osInfo["version"],
		Build:      osInfo["build"],
		OsqueryVer: osqueryVer["version"],
		Timestamp:  time.Now(),
	}

	if err := tx.Create(&newOS).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert Applications
	for _, app := range apps {
		application := Application{
			ID:      app["id"],
			Name:    app["name"],
			Version: app["version"],
			Path:    app["path"],
		}
		if err := tx.Create(&application).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	return tx.Commit().Error
}

// FetchLatestSnapshot retrieves the most recent system snapshot
func FetchLatestSnapshot() (OSInfo, []Application, error) {

	// Get latest OS info
	var latestOS OSInfo
	err := db.Order("timestamp desc").First(&latestOS).Error
	if err != nil {
		return latestOS, nil, err
	}

	// Get associated applications
	var appRecords []Application
	err = db.Where("os_info_id = ?", latestOS.ID).Find(&appRecords).Error
	if err != nil {
		return latestOS, nil, err
	}

	return latestOS, appRecords, nil
}

func getDatabaseURL() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
}
