package handlers

import (
	"antreanPuskesmas/config"
	"antreanPuskesmas/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "Selamat datang di Sistem Antrean Puskesmas")
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "pasien_register.html", nil)
}

func RegisterPasien(c *gin.Context) {
	var input struct {
		NamaLengkap        string `form:"nama" binding:"required"`
		Email              string `form:"email" binding:"required"`
		Password           string `form:"password" binding:"required"`
		KonfirmasiPassword string `form:"konfirmasi_password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak valid: %v", err)
		return
	}

	if input.Password != input.KonfirmasiPassword {
		c.String(http.StatusBadRequest, "Password dan konfirmasi tidak sama")
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Simpan ke database
	pasien := models.Pasien{
		NamaLengkap: input.NamaLengkap,
		Email:       input.Email,
		Password:    string(hashedPassword),
	}

	if err := config.DB.Create(&pasien).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan data pasien")
		return
	}

	c.String(http.StatusOK, "Registrasi berhasil! Silakan login.")
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "pasien_login.html", nil)
}

func LoginPasien(c *gin.Context) {
	var input struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Input tidak valid")
		return
	}

	var pasien models.Pasien
	err := config.DB.Where("email = ?", input.Email).First(&pasien).Error
	if err != nil {
		c.String(http.StatusUnauthorized, "Email tidak ditemukan")
		return
	}

	// Cek password
	err = bcrypt.CompareHashAndPassword([]byte(pasien.Password), []byte(input.Password))
	if err != nil {
		c.String(http.StatusUnauthorized, "Password salah")
		return
	}

	// Simpan sesi login
	session := sessions.Default(c)
	session.Set("user_id", pasien.ID)
	session.Set("user_name", pasien.NamaLengkap)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func PasienDashboard(c *gin.Context) {
	session := sessions.Default(c)
	nama := session.Get("user_name")
	if nama == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "pasien_dashboard.html", gin.H{
		"Nama": nama,
	})
}

func LogoutPasien(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
