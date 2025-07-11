package routes

import (
	"antreanPuskesmas/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Halaman utama
	r.GET("/", handlers.HomePage)

	// Pasien
	r.GET("/register", handlers.ShowRegisterPage)
	r.POST("/register", handlers.RegisterPasien)
	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.LoginPasien)
	r.GET("/dashboard", handlers.PasienDashboard)

	// Antrean
	r.GET("/antrean/daftar", handlers.ShowFormAntrean)
	r.POST("/antrean/daftar", handlers.SubmitAntrean)
	r.POST("/antrean/preview", handlers.PreviewAntrean)
	r.POST("/antrean/simpan", handlers.SimpanAntrean)
	r.GET("/antrean/status", handlers.StatusAntrean)
	r.GET("/antrean/status/data", handlers.GetStatusData)
	r.GET("/antrean/riwayat", handlers.RiwayatPasien)

	// Petugas
	r.GET("/petugas/login", handlers.ShowPetugasLogin)
	r.POST("/petugas/login", handlers.LoginPetugas)
	r.GET("/petugas/dashboard", handlers.PetugasDashboard)
	r.GET("/petugas/antrean", handlers.PetugasAntreanPage)

	//
	r.POST("/petugas/tambah", handlers.TambahPasienManual)
	r.POST("/petugas/update/:id", handlers.UpdateAntrean)
	r.POST("/petugas/delete/:id", handlers.HapusAntrean)
	r.POST("/petugas/panggil/:id", handlers.PanggilPasien)
	r.POST("/petugas/selesai/:id", handlers.SelesaiAntrean)

	r.POST("/petugas/riwayat/simpan", handlers.SimpanKeRiwayat)
	r.GET("/petugas/riwayat", handlers.RiwayatBulananPage)
	r.GET("/petugas/riwayat/:tahun/:bulan", handlers.RiwayatHarianPage)
	r.GET("/petugas/riwayat/:tahun/:bulan/cetak", handlers.CetakRiwayatPDF)
	r.GET("/petugas/riwayat/:tahun/:bulan/hapus", handlers.HapusRiwayatBulanan)

	r.GET("/logout", handlers.LogoutPasien)
	r.GET("/petugas/logout", handlers.LogoutPetugas)

}
