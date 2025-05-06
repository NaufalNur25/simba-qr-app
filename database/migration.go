package database

import (
	"fmt"

	"github.com/naufal/simba-qr-app/config"
	"github.com/naufal/simba-qr-app/models"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&models.System{},
	)

	if err != nil {
		panic("❌ Migration Failed: " + err.Error())
	}

	fmt.Println("✅ Database Migrated Successfully")
}
