package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=ep-wispy-darkness-a1munvp0-pooler.ap-southeast-1.aws.neon.tech user=neondb_owner password=npg_JKyp46IsbBia dbname=neondb port=5432 sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi gagal karena : ", err)
	}

	DB = db
	fmt.Println("koneksi sukses")
}
