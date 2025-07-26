// File: models/item.go
// Deskripsi: Mendefinisikan struktur data (model) untuk 'Item'.

package models

// Item mendefinisikan struktur untuk entitas buku.
// Nama field harus diawali huruf kapital agar bisa diakses dari package lain.
// Tag `json:"..."` digunakan untuk mengontrol bagaimana field ini di-encode/decode ke/dari JSON.
// Tag `binding:"required"` adalah untuk validasi oleh Gin saat binding data.
type Item struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required" Gorm:"unique"`
	Author string `json:"author" binding:"required"`
}
