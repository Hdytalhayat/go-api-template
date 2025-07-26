// File: main.go
// Deskripsi: Titik masuk utama (entrypoint) untuk seluruh aplikasi.
// Tanggung jawabnya adalah:
// 1. Memuat konfigurasi dari file .env.
// 2. Menginisialisasi koneksi database.
// 3. Membuat instance Gin router.
// 4. Mendaftarkan semua rute API.
// 5. Menjalankan server HTTP.

package main

import (
	"fmt"
	"log"
	"os"

	"go-api-template/config"
	"go-api-template/routes"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Memuat konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak dapat menemukan file .env, menggunakan variabel lingkungan sistem.")
	}

	// 2. Menginisialisasi koneksi database
	config.ConnectDatabase()
	// Menjadwalkan penutupan koneksi database saat fungsi main selesai
	defer config.DB.Close()

	// 3. Membuat dan mengatur router
	// Fungsi SetupRouter dari package routes akan mengembalikan instance Gin yang sudah dikonfigurasi
	router := routes.SetupRouter()

	// 4. Menjalankan server
	// Mengambil port dari environment variable, dengan fallback ke 8080
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port :%s", port)
	// Menjalankan server di alamat dan port yang telah ditentukan
	router.Run(fmt.Sprintf(":%s", port))
}
