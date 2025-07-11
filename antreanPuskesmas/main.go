package main

import (
	"antreanPuskesmas/config"
	"antreanPuskesmas/routes"
	"html/template"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	// Session
	store := cookie.NewStore([]byte("secret123"))
	r.Use(sessions.Sessions("mysession", store))

	// Template function map
	funcMap := template.FuncMap{
		"title": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToUpper(s[:1]) + s[1:]
		},
		"add": func(a, b int) int {
			return a + b
		},
		"bulanNama": func(bulan int) string {
			bulanMap := [...]string{
				"", "JANUARI", "FEBRUARI", "MARET", "APRIL", "MEI", "JUNI",
				"JULI", "AGUSTUS", "SEPTEMBER", "OKTOBER", "NOVEMBER", "DESEMBER",
			}
			if bulan >= 1 && bulan <= 12 {
				return bulanMap[bulan]
			}
			return "Tidak valid"
		},
	}

	// Load template with custom functions
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*"))
	r.SetHTMLTemplate(tmpl)

	// Static files and routing
	r.Static("/static", "./static")
	routes.SetupRoutes(r)

	r.Run(":8080")
}
