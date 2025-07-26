# Template REST API Go (Gin)

Ini adalah templat dasar yang kokoh untuk membangun REST API menggunakan Go dengan framework Gin.

## Fitur

-   Struktur Proyek yang Rapi (MVC-like)
-   Manajemen Konfigurasi via file `.env`
-   Routing Terpusat
-   Koneksi Database MySQL
-   Helper untuk Respons JSON Standar

## Prasyarat

-   Go versi 1.18 atau lebih baru
-   Server Database MySQL yang sedang berjalan

## Cara Menggunakan

1.  **Kloning Repositori**
    ```bash
    git clone https://github.com/Hdytalhayat/go-api-template.git nama-proyek-baru
    cd nama-proyek-baru
    ```

2.  **Ubah Nama Modul (Opsional tapi Direkomendasikan)**
    Buka file `go.mod` dan ubah `go-api-template` menjadi nama proyek Anda, misalnya `my-awesome-api`. Lakukan juga find & replace di seluruh proyek.

3.  **Konfigurasi Lingkungan**
    Salin file `.env.example` menjadi `.env` dan sesuaikan nilainya dengan konfigurasi database Anda.
    ```bash
    cp .env.example .env
    ```

4.  **Install Dependensi**
    ```bash
    go mod tidy
    ```

5.  **Jalankan Server**
    ```bash
    go run main.go
    ```
    Server akan berjalan di port yang ditentukan di file `.env` (default: 8080).

## Struktur Proyek

-   `main.go`: Titik masuk utama aplikasi.
-   `/config`: Koneksi dan konfigurasi (database, dll).
-   `/controllers`: Berisi logika untuk menangani request HTTP.
-   `/models`: Definisi struct data (misal: User, Product).
-   `/routes`: Definisi semua endpoint API.
-   `/utils`: Fungsi helper yang dapat digunakan kembali.
