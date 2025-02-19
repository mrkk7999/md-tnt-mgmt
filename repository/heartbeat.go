package repository

import (
	"log"
)

// HeartBeat
func (r *repository) HeartBeat() map[string]string {
	sqlDB, err := r.db.DB()
	if err != nil {
		log.Println("Error getting DB instance:", err)
		return map[string]string{
			"message": "Service is running, but DB check failed",
			"db":      "Error retrieving DB instance",
		}
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Database Ping Failed:", err)
		return map[string]string{
			"message": "Service is running, but DB is unreachable",
			"db":      "Down",
		}
	}

	return map[string]string{
		"message": "Service is up and running",
		"db":      "Up",
	}
}
