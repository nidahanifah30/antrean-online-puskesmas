package config

import (
	"antreanPuskesmas/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/antrean_puskesmas?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	fmt.Println("Koneksi database berhasil!")

	// Auto migrate semua model
	err = DB.AutoMigrate(
		&models.Pasien{},
		&models.Petugas{},
		&models.Antrean{},
		&models.RiwayatHarian{},
		&models.RiwayatBulanan{},
	)
	if err != nil {
		log.Fatal("AutoMigrate gagal:", err)
	}

	fmt.Println("Migrasi database berhasil!")
}
