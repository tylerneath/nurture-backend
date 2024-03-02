package store

import (
	"log"
	"time"

	"github.com/tylerneath/nuture-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustCreateNewDB(cfg config.Config) *gorm.DB {

	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i < 20; i++ {
		db, err = gorm.Open(postgres.Open(cfg.Dsn()), &gorm.Config{
			// Logger:         logger.Default.LogMode(logger.Silent),
			TranslateError: true,
		})
		if err != nil {
			log.Printf("Failed to connect to db, retrying... %v", err)
			time.Sleep(2 * time.Second)
		} else {
			log.Printf("Successfully connected to db")
			break

		}
	}
	if err != nil {
		log.Fatalf("Fatal error connecting to db: %v", err)
	}

	return db

}
