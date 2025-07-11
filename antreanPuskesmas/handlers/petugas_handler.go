package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"antreanPuskesmas/config"
	"antreanPuskesmas/models"

	"github.com/phpdave11/gofpdf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowPetugasLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "petugas_login.html", nil)
}

func LoginPetugas(c *gin.Context) {
	var input struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	c.ShouldBind(&input)

	var petugas models.Petugas
	err := config.DB.Where("username = ?", input.Username).First(&petugas).Error
	if err != nil {
		c.String(http.StatusUnauthorized, "Username tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(petugas.Password), []byte(input.Password)); err != nil {
		c.String(http.StatusUnauthorized, "Password salah")
		return
	}

	session := sessions.Default(c)
	session.Set("petugas_id", petugas.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/petugas/dashboard")
}

func PetugasDashboard(c *gin.Context) {
	today := time.Now().Format("2006-01-02")

	var total, menunggu, selesai int64

	// Total semua antrean hari ini
	config.DB.Model(&models.Antrean{}).
		Where("DATE(tanggal_kunjungan) = ?", today).
		Count(&total)

	// Jumlah menunggu
	config.DB.Model(&models.Antrean{}).
		Where("DATE(tanggal_kunjungan) = ? AND status = ?", today, "menunggu").
		Count(&menunggu)

	// Jumlah selesai
	config.DB.Model(&models.Antrean{}).
		Where("DATE(tanggal_kunjungan) = ? AND status = ?", today, "selesai").
		Count(&selesai)

	c.HTML(http.StatusOK, "petugas_dashboard.html", gin.H{
		"Total":    total,
		"Menunggu": menunggu,
		"Selesai":  selesai,
	})
}
func PetugasAntreanPage(c *gin.Context) {
	today := time.Now().Format("2006-01-02")
	var antrean []models.Antrean

	config.DB.
		Where("DATE(tanggal_kunjungan) = ?", today).
		Order("created_at").
		Find(&antrean)

	c.HTML(http.StatusOK, "petugas_antrean.html", gin.H{
		"Antrean": antrean,
		"Now":     time.Now().Format("2006-01-02"),
	})
}

func TambahPasienManual(c *gin.Context) {
	var input struct {
		Nama    string `form:"nama"`
		NIK     string `form:"nik"`
		NoHP    string `form:"no_hp"`
		Layanan string `form:"layanan"`
		Tanggal string `form:"tanggal"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	fmt.Printf("Input: %+v\n", input)

	// Parse tanggal
	t, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, "Tanggal tidak valid")
		return
	}

	// Hitung antrean hari ini
	var count int64
	config.DB.Model(&models.Antrean{}).Where("tanggal_kunjungan = ?", t).Count(&count)
	nomor := fmt.Sprintf("A%03d", count+1)

	newAntrean := models.Antrean{
		NamaLengkap:      input.Nama,
		NIK:              input.NIK,
		NoHP:             input.NoHP,
		Layanan:          input.Layanan,
		TanggalKunjungan: t,
		NomorAntrean:     nomor,
		Status:           "menunggu",
	}

	if err := config.DB.Create(&newAntrean).Error; err != nil {
		fmt.Println("DB error:", err)
		c.String(http.StatusInternalServerError, "Gagal menyimpan antrean: "+err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/petugas/antrean")
}

func UpdateAntrean(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Nama    string `form:"nama"`
		NIK     string `form:"nik"`
		NoHP    string `form:"no_hp"`
		Layanan string `form:"layanan"`
		Tanggal string `form:"tanggal"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak valid")
		return
	}

	tanggal, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		c.String(http.StatusBadRequest, "Tanggal tidak valid")
		return
	}

	config.DB.Model(&models.Antrean{}).Where("id = ?", id).Updates(models.Antrean{
		NamaLengkap:      input.Nama,
		NIK:              input.NIK,
		NoHP:             input.NoHP,
		Layanan:          input.Layanan,
		TanggalKunjungan: tanggal,
	})

	c.Redirect(http.StatusFound, "/petugas/antrean")
}

func HapusAntrean(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Antrean{}, id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus antrean")
		return
	}

	c.Redirect(http.StatusFound, "/petugas/antrean")
}

func PanggilPasien(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Model(&models.Antrean{}).
		Where("id = ?", id).
		Update("status", "dipanggil").Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal memanggil pasien")
		return
	}

	c.Redirect(http.StatusFound, "/petugas/antrean")
}

func SelesaiAntrean(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Model(&models.Antrean{}).
		Where("id = ?", id).
		Update("status", "selesai").Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyelesaikan antrean")
		return
	}

	c.Redirect(http.StatusFound, "/petugas/antrean")
}

func SimpanKeRiwayat(c *gin.Context) {
	today := time.Now().Format("2006-01-02")

	// Ambil semua antrean selesai hari ini
	var antreanSelesai []models.Antrean
	config.DB.
		Where("DATE(tanggal_kunjungan) = ? AND status = ?", today, "selesai").
		Find(&antreanSelesai)

	// Simpan ke Riwayat Harian
	for _, a := range antreanSelesai {
		r := models.RiwayatHarian{
			Tanggal:      a.TanggalKunjungan,
			NomorAntrean: a.NomorAntrean,
			NamaPasien:   a.NamaLengkap,
			Layanan:      a.Layanan,
			Status:       "selesai",
		}
		config.DB.Create(&r)
	}

	// Hitung jumlah pasien untuk bulan ini
	now := time.Now()
	tahun := now.Year()
	bulan := int(now.Month())

	var totalPasien int64
	config.DB.Model(&models.RiwayatHarian{}).
		Where("YEAR(tanggal) = ? AND MONTH(tanggal) = ?", tahun, bulan).
		Count(&totalPasien)

	// Cek jika sudah ada
	var existing models.RiwayatBulanan
	if err := config.DB.
		Where("tahun = ? AND bulan = ?", tahun, bulan).
		First(&existing).Error; err == nil {
		// Update jumlah
		existing.JumlahPasien = int(totalPasien)
		config.DB.Save(&existing)
	} else {
		// Tambah baru
		config.DB.Create(&models.RiwayatBulanan{
			Tahun:        tahun,
			Bulan:        bulan,
			JumlahPasien: int(totalPasien),
		})
	}

	c.Redirect(http.StatusFound, "/petugas/riwayat")
}

func RiwayatBulananPage(c *gin.Context) {
	var data []models.RiwayatBulanan
	config.DB.Order("tahun desc, bulan desc").Find(&data)

	c.HTML(http.StatusOK, "riwayat_bulanan.html", gin.H{
		"Data": data,
	})
}

func RiwayatHarianPage(c *gin.Context) {
	tahunStr := c.Param("tahun")
	bulanStr := c.Param("bulan")

	tahunInt, _ := strconv.Atoi(tahunStr)
	bulanInt, _ := strconv.Atoi(bulanStr)

	var riwayat []models.RiwayatHarian
	config.DB.
		Where("YEAR(tanggal) = ? AND MONTH(tanggal) = ?", tahunStr, bulanStr).
		Order("tanggal").
		Find(&riwayat)

	c.HTML(http.StatusOK, "riwayat_harian.html", gin.H{
		"Riwayat": riwayat,
		"Tahun":   tahunInt,
		"Bulan":   bulanInt,
	})
}

func CetakRiwayatPDF(c *gin.Context) {
	tahun := c.Param("tahun")
	bulan := c.Param("bulan")

	// Ambil data riwayat
	var riwayat []models.RiwayatHarian
	config.DB.
		Where("YEAR(tanggal) = ? AND MONTH(tanggal) = ?", tahun, bulan).
		Order("tanggal").
		Find(&riwayat)

	// Konversi bulan ke nama
	bulanNama := getNamaBulan(bulan)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Header dan Judul
	pdf.SetFillColor(0, 0, 0)      // hitam
	pdf.Rect(10, 10, 190, 30, "F") // header block

	// Judul di tengah
	pdf.SetXY(50, 15)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(110, 7, "PUSKESMAS KETANGGUNGAN", "", 1, "C", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(110, 7, "LAPORAN PELAYANAN PASIEN", "", 1, "C", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(110, 7, fmt.Sprintf("BULAN %s %s", bulanNama, tahun), "", 1, "C", false, 0, "")

	pdf.Ln(15)   // Jarak dari header ke tabel
	pdf.SetX(15) // Posisi tabel ke tengah kanan

	// Tabel Header
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(70, 70, 70)
	pdf.SetTextColor(255, 255, 255)

	pdf.SetX(15)
	pdf.CellFormat(10, 8, "NO", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "TANGGAL", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 8, "NOMOR ANTREAN", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 8, "NAMA PASIEN", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "LAYANAN", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "STATUS", "1", 1, "C", true, 0, "")

	// Tabel Body
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)
	for i, r := range riwayat {
		pdf.SetX(15)
		pdf.CellFormat(10, 8, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, r.Tanggal.Format("02-01-2006"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 8, r.NomorAntrean, "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 8, r.NamaPasien, "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 8, r.Layanan, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, r.Status, "1", 1, "C", false, 0, "")
	}

	pdf.Ln(5)

	// Footer info
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(50, 8, "Total Pasien :", "", 0, "", false, 0, "")
	pdf.CellFormat(0, 8, fmt.Sprintf("%d", len(riwayat)), "", 1, "", false, 0, "")
	pdf.CellFormat(50, 8, "Dicetak oleh :", "", 0, "", false, 0, "")
	pdf.CellFormat(0, 8, "PETUGAS", "", 1, "", false, 0, "")
	pdf.CellFormat(50, 8, "Tanggal Cetak :", "", 0, "", false, 0, "")
	pdf.CellFormat(0, 8, time.Now().Format("02-01-2006"), "", 1, "", false, 0, "")

	// Tanda tangan di kanan
	pdf.Ln(20)
	pdf.CellFormat(0, 8, "Tanda Tangan Petugas", "", 1, "R", false, 0, "")
	pdf.Ln(20)
	pdf.CellFormat(0, 8, "(.................................)", "", 1, "R", false, 0, "")

	filename := fmt.Sprintf("riwayat_%s_%s.pdf", tahun, bulan)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	_ = pdf.Output(c.Writer)
}

func getNamaBulan(bulan string) string {
	bulanInt, err := strconv.Atoi(bulan)
	if err != nil || bulanInt < 1 || bulanInt > 12 {
		return "Tidak Valid"
	}
	daftar := []string{
		"", "JANUARI", "FEBRUARI", "MARET", "APRIL", "MEI", "JUNI",
		"JULI", "AGUSTUS", "SEPTEMBER", "OKTOBER", "NOVEMBER", "DESEMBER",
	}
	return daftar[bulanInt]
}

func HapusRiwayatBulanan(c *gin.Context) {
	tahun := c.Param("tahun")
	bulan := c.Param("bulan")

	// Hapus semua riwayat harian pada bulan & tahun tersebut
	if err := config.DB.
		Where("YEAR(tanggal) = ? AND MONTH(tanggal) = ?", tahun, bulan).
		Delete(&models.RiwayatHarian{}).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus riwayat harian")
		return
	}

	// Hapus entri bulanan
	if err := config.DB.
		Where("tahun = ? AND bulan = ?", tahun, bulan).
		Delete(&models.RiwayatBulanan{}).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus riwayat bulanan")
		return
	}

	c.Redirect(http.StatusFound, "/petugas/riwayat")
}

func LogoutPetugas(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/petugas/login")
}
