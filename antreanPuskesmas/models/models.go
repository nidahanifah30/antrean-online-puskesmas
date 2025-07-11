package models

import (
	"time"
)

type Pasien struct {
	ID          uint   `gorm:"primaryKey"`
	NamaLengkap string `gorm:"size:100;not null"`
	Email       string `gorm:"size:100;unique;not null"`
	Password    string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	Antrean     []Antrean `gorm:"foreignKey:PasienID"`
}

type Petugas struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:50;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
}

type Antrean struct {
	ID               uint   `gorm:"primaryKey"`
	PasienID         *uint  `gorm:"default:null"` // ← ✅ pakai pointer & nullable
	NamaLengkap      string `gorm:"size:100;not null"`
	NIK              string `gorm:"size:20;not null"`
	NoHP             string `gorm:"size:20;not null"`
	Layanan          string `gorm:"size:100;not null"`
	TanggalKunjungan time.Time
	NomorAntrean     string `gorm:"size:10;not null"`
	Status           string `gorm:"type:enum('menunggu','dipanggil','selesai');default:'menunggu'"`
	CreatedAt        time.Time
}

type RiwayatHarian struct {
	ID           uint      `gorm:"primaryKey"`
	Tanggal      time.Time `gorm:"not null"`
	NomorAntrean string    `gorm:"size:10;not null"`
	NamaPasien   string    `gorm:"size:100;not null"`
	Layanan      string    `gorm:"size:100;not null"`
	Status       string    `gorm:"type:enum('selesai');default:'selesai'"`
	CreatedAt    time.Time
}

type RiwayatBulanan struct {
	ID           uint `gorm:"primaryKey"`
	Tahun        int  `gorm:"not null"`
	Bulan        int  `gorm:"not null"`
	JumlahPasien int  `gorm:"not null"`
	CreatedAt    time.Time
}
