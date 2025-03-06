package database

import (
	"log"

	eu "go-commerce-api/internal/user/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	migrator := db.Migrator()

	db.AutoMigrate(
		&eu.User{},
	)

	tables := []string{"users", "products", "payments"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
