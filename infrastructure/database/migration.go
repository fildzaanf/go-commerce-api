package database

import (
	"log"

	eu "go-commerce-api/internal/user/entity"
	epy "go-commerce-api/internal/payment/entity"
	epd "go-commerce-api/internal/product/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	migrator := db.Migrator()

	db.AutoMigrate(
		&eu.User{},
		&epd.Product{},
		&epy.Payment{},
	)

	tables := []string{"users", "products", "payments"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
