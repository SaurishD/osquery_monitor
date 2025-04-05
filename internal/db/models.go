// internal/db/models.go
package db

import "time"

// OSInfo represents a snapshot of the operating system information
type OSInfo struct {
	ID         int       `json:"id"`
	Platform   string    `json:"platform"`
	Version    string    `json:"version"`
	Build      string    `json:"build"`
	OsqueryVer string    `json:"osquery_version"`
	Timestamp  time.Time `json:"timestamp"`
}

// Application represents an installed application
type Application struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	OSInfoID int    `json:"os_info_id"`
}
