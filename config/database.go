// File: config/database.go
// Deskripsi: Mengelola koneksi ke database.

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah variabel global yang akan menampung instance koneksi database.
// Variabel ini diekspor (huruf kapital) agar bisa diakses dari package lain.
var DB *sql.DB

// ConnectDatabase adalah fungsi untuk membuat koneksi ke database MySQL.
// Fungsi ini membaca konfigurasi dari environment variables.
func ConnectDatabase() {
	// Membaca konfigurasi dari environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Membuat Data Source Name (DSN) string
	// format: "user:password@tcp(host:port)/dbname?parseTime=true"
	// parseTime=true penting agar Go bisa mem-parsing tipe data DATETIME dan TIMESTAMP dari MySQL.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	// Membuka koneksi ke database dan menyimpannya ke variabel global DB
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Fatal Error: Gagal membuka koneksi ke database: %v", err)
	}

	// Memeriksa apakah koneksi benar-benar berhasil dibuat
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Fatal Error: Gagal melakukan ping ke database: %v", err)
	}

	log.Println("Koneksi ke database berhasil dibuat.")
}
