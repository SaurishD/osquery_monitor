package db

import "time"

type OSInfo struct {
	ID         int       `json:"id"`
	Platform   string    `json:"platform"`
	Version    string    `json:"version"`
	Build      string    `json:"build"`
	OsqueryVer string    `json:"osquery_version"`
	Timestamp  time.Time `json:"timestamp"`
}

type Application struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Path    string `json:"path" gorm:"primaryKey"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
