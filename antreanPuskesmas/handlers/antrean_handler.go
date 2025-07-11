package handlers

import (
	"fmt"
	"net/http"
	"time"

	"antreanPuskesmas/config"
	"antreanPuskesmas/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Form input pendaftaran antrean
func ShowFormAntrean(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("user_name")
	if nama == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "antrean_form.html", gin.H{
		"Nama": nama,
	})
}

func SubmitAntrean(c *gin.Context) {
	session := sessions.Default(c)
	pasienID := session.Get("user_id")
	nama := session.Get("user_name")

	if pasienID == nil || nama == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var input struct {
		NIK     string `form:"nik" binding:"required"`
		NoHP    string `form:"no_hp" binding:"required"`
		Layanan string `form:"layanan" binding:"required"`
		Tanggal string `form:"tanggal" binding:"required"` // format: YYYY-MM-DD
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak lengkap")
		return
	}

	// Generate nomor antrean: contoh A001, A002, ...
	var count int64
	config.DB.Model(&models.Antrean{}).Where("tanggal_kunjungan = ?", input.Tanggal).Count(&count)
	nomor := fmt.Sprintf("A%03d", count+1)

	// Simpan antrean
	data := models.Antrean{
		PasienID:         ptrUint(pasienID.(uint)),
		NamaLengkap:      nama.(string),
		NIK:              input.NIK,
		NoHP:             input.NoHP,
		Layanan:          input.Layanan,
		TanggalKunjungan: parseDate(input.Tanggal),
		NomorAntrean:     nomor,
		Status:           "menunggu",
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal simpan antrean")
		return
	}

	// Tampilkan nomor
	c.HTML(http.StatusOK, "antrean_sukses.html", gin.H{
		"Nomor":   nomor,
		"Nama":    nama,
		"Layanan": input.Layanan,
		"Tanggal": input.Tanggal,
	})
}

func ptrUint(v uint) *uint {
	return &v
}

func parseDate(str string) time.Time {
	t, _ := time.Parse("2006-01-02", str)
	return t
}

func PreviewAntrean(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("user_name")
	pasienID := session.Get("user_id")

	var input struct {
		NIK     string `form:"nik"`
		NoHP    string `form:"no_hp"`
		Layanan string `form:"layanan"`
		Tanggal string `form:"tanggal"`
	}

	c.ShouldBind(&input)

	// Hitung nomor antrean
	var count int64
	config.DB.Model(&models.Antrean{}).Where("tanggal_kunjungan = ?", input.Tanggal).Count(&count)
	nomor := fmt.Sprintf("A%03d", count+1)

	c.JSON(200, gin.H{
		"pasien_id": pasienID,
		"nama":      nama,
		"nik":       input.NIK,
		"no_hp":     input.NoHP,
		"layanan":   input.Layanan,
		"tanggal":   input.Tanggal,
		"nomor":     nomor,
	})
}

func SimpanAntrean(c *gin.Context) {
	session := sessions.Default(c)
	pasienID := session.Get("user_id")

	if pasienID == nil {
		c.String(http.StatusUnauthorized, "Silakan login dulu")
		return
	}

	var input struct {
		Nama    string `form:"nama"`
		NIK     string `form:"nik"`
		NoHP    string `form:"no_hp"`
		Layanan string `form:"layanan"`
		Tanggal string `form:"tanggal"`
		Nomor   string `form:"nomor"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak lengkap")
		return
	}

	data := models.Antrean{
		PasienID:         ptrUint(pasienID.(uint)),
		NamaLengkap:      input.Nama,
		NIK:              input.NIK,
		NoHP:             input.NoHP,
		Layanan:          input.Layanan,
		TanggalKunjungan: parseDate(input.Tanggal),
		NomorAntrean:     input.Nomor,
		Status:           "menunggu",
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal simpan antrean: "+err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func StatusAntrean(c *gin.Context) {
	session := sessions.Default(c)
	pasienID := session.Get("user_id")
	nama := session.Get("user_name")

	if pasienID == nil || nama == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var antrean models.Antrean
	err := config.DB.
		Where("pasien_id = ?", pasienID).
		Order("created_at desc").
		First(&antrean).Error

	if err != nil {
		c.HTML(http.StatusOK, "status_antrean.html", gin.H{
			"AdaAntrean": false,
			"Nama":       nama,
		})
		return
	}

	// Hitung jumlah antrean yang statusnya "menunggu" dengan tanggal yang sama dan created_at lebih awal
	var antreanSebelum int64
	config.DB.Model(&models.Antrean{}).
		Where("tanggal_kunjungan = ? AND status IN ? AND created_at < ?", antrean.TanggalKunjungan, []string{"menunggu", "dipanggil"}, antrean.CreatedAt).
		Count(&antreanSebelum)

	estimasiMenit := antreanSebelum * 5 // misal 1 antrean = 5 menit

	c.HTML(http.StatusOK, "status_antrean.html", gin.H{
		"AdaAntrean":    true,
		"Nama":          nama,
		"Antrean":       antrean,
		"EstimasiMenit": estimasiMenit,
	})
}

func GetStatusData(c *gin.Context) {
	session := sessions.Default(c)
	pasienID := session.Get("user_id")

	var antrean models.Antrean
	err := config.DB.
		Where("pasien_id = ?", pasienID).
		Order("created_at desc").
		First(&antrean).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Antrean tidak ditemukan"})
		return
	}

	var antreanSebelum int64
	config.DB.Model(&models.Antrean{}).
		Where("tanggal_kunjungan = ? AND status IN ? AND created_at < ?", antrean.TanggalKunjungan, []string{"menunggu", "dipanggil"}, antrean.CreatedAt).
		Count(&antreanSebelum)

	c.JSON(200, gin.H{
		"status":        antrean.Status,
		"nomor":         antrean.NomorAntrean,
		"estimasiMenit": antreanSebelum * 5,
	})
}

func RiwayatPasien(c *gin.Context) {
	session := sessions.Default(c)
	pasienID := session.Get("user_id")
	nama := session.Get("user_name")

	if pasienID == nil || nama == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var riwayat []models.Antrean
	config.DB.
		Where("pasien_id = ?", pasienID).
		Order("created_at DESC").
		Find(&riwayat)

	c.HTML(http.StatusOK, "riwayat_antrean.html", gin.H{
		"Nama":    nama,
		"Riwayat": riwayat,
	})
}
