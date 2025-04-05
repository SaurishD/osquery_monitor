package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Init initializes the database connection
func Init() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/osquery_monitor?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	return db
}

func InsertLatestSnapshot(db *sql.DB, osInfo map[string]string, osqueryVer map[string]string, apps []map[string]string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert OS Info
	var osInfoID int
	err = tx.QueryRow(`INSERT INTO os_info (platform, version, build, osquery_version, timestamp)
        VALUES ($1, $2, $3, $4, NOW()) RETURNING id`,
		osInfo["platform"], osInfo["version"], osInfo["build"], osqueryVer["version"]).Scan(&osInfoID)
	if err != nil {
		return err
	}

	// Insert Applications
	for _, app := range apps {
		_, err = tx.Exec(`INSERT INTO applications (name, version, os_info_id)
            VALUES ($1, $2, $3)`, app["name"], app["version"], osInfoID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// FetchLatestSnapshot retrieves the most recent system snapshot
func FetchLatestSnapshot() (map[string]string, []map[string]string, error) {
	db := Init()
	defer db.Close()

	// Get latest OS info
	osInfo := make(map[string]string)
	var platform, version, build, osqueryVersion string
	err := db.QueryRow(`
		SELECT platform, version, build, osquery_version 
		FROM os_info 
		ORDER BY timestamp DESC 
		LIMIT 1
	`).Scan(&platform, &version, &build, &osqueryVersion)
	if err != nil {
		return nil, nil, err
	}

	osInfo["platform"] = platform
	osInfo["version"] = version
	osInfo["build"] = build
	osInfo["osquery_version"] = osqueryVersion

	// Get associated applications
	rows, err := db.Query(`
		SELECT name, version 
		FROM applications 
		WHERE os_info_id = (
			SELECT id 
			FROM os_info 
			ORDER BY timestamp DESC 
			LIMIT 1
		)
	`)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var apps []map[string]string
	for rows.Next() {
		var name, version string
		err := rows.Scan(&name, &version)
		if err != nil {
			return nil, nil, err
		}
		app := map[string]string{
			"name":    name,
			"version": version,
		}
		apps = append(apps, app)
	}

	return osInfo, apps, nil
}
