package main // main untuk menandakan bahwa ini adalah package utama yang akan dijalankan.
// package nama lain untuk menandakan package libary, misal package db, package utils, package models, dll.

import (
	// dibawah ini adalah modul/library yang akan digunakan di dalam project ini

	// ini adalah modul bawaan dari golang, jadi tidak perlu di install lagi
	"context" // untuk mengatur rem darurat, misal timeout, bisa juga digunakan untuk promise atau async (menunggu proses selesai)
	"fmt"     // untuk print console
	"log"     // mirip print console, tapi bisa menampilkan lebih lengkap, misal error
	"os"      // untuk mengakses file system, misal membaca file, men

	// ini adalah modul/library dari internet atau modul orang lain, menggunakan perintah go get <nama modul>
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	// ini adalah modul/library yang dibuat sendiri, jadi harus di import sesuai dengan path projectnya
	"belajar-api/internal/db"
)

// fungsi main adalah fungsi utama yang akan dijalankan pertama kali ketika program dijalankan.
// fungsi main wajib ada jika package main, tapi tidak wajib jika package lain.
func main() {
	err := godotenv.Load()
	// err adalah variabel yang akan menampung error jika terjadi error saat menajalankan fungsi godotenv.Load().
	// godotenv ini adalah nama library/modul yang di panggil
	// Load() adalah fungsi dari library godotenv yang digunakan untuk membaca file .env
	if err != nil {
	// variabel err tidak nil/null/kosong maka akan menjalankan dibawah ini.
	// jadi artinya jika output dari godotenv.Load() null atau kosong, program berhasil berjalan. tapi jika ada output/isinya berarti program ada yang error
		log.Fatalf("Gagal load file .env: %v\n", err)
		// log adalah modul/library yang di import dari log, digunakan untuk menampilkan log error di console.
		// Fatalf berfungsi mencetak/print sesuai format yang ditentukan, lalu menghentikan program.
		// %v ini adala tempat kosong yang akan di gantikan dengan nilai yang akan diberikan setelah koma, dalam hal ini adalah err.
	}

	dbURL := os.Getenv("GOOSE_DBSTRING")
	// dbURL adalah variabel yang akan menampung nilai dari os.Getenv("GOOSE_DBSTRING")
	// os adalah modul/library yang di import dari os, digunakan untuk mengakses file system, misal membaca file, menulis file, dll.
	// Getenv adalah fungsi dari modul os yang digunakan untuk membaca nilai dari environment variable, dalam hal ini adalah GOOSE_DBSTRING.
	if dbURL == "" {
	// jika dbURL kosong, maka akan menjalankan dibawah ini.
		log.Fatal("GOOSE_DBSTRING tidak ditemukan di file .env")
		// log.Fatal akan menampilkan log error dan menghentikan program.
		// disini tidak menggunakan Fatalf karena tidak ada format tertentu yang akan ditampilkan yang cukup string saja.
	}

	ctx := context.Background()
	// variabel ctx untuk menampung context.Background()
	// context adalah modul/library yang di import dari "context", digunakan untuk mengatur rem darurat.
	// Background() adalah fungsi dari modul context yang digunakan untuk menunggu proses sampai selesai, bisa juga digunakan untuk promise atau async (menunggu proses sampai selesai).
	pool, err := pgxpool.New(ctx, dbURL)
	// pgxpool.New() adalah fungsi untuk membuat koneksi ke databse PostgreSQL, lalu alamat koneksinya di simpan ke variabel pool.
	// disini mendeklarasikan 2 variabel, pool dan err.
	// disini variabel err akan menggantikan atau menghapus varibel err sebelumnya, karena menggunakan tanda (:=) yang artinya mendeklarasikan variabel baru.
	// pgxpool adalah modul/library yang di import dari "github.com/jackc/pgx/v5/pgxpool", digunakan untuk mengatur koneksi ke database PostgreSQL.
	// New() adalah fungsi dari modul pgxpool yang digunakan untuk membuat koneksi baru ke database PostgreSQL.
	if err != nil {
		log.Fatalf("Gagal connect ke database: %v\n", err)
	}
	defer pool.Close()
	// defer adalah fungsi yang akan dijalankan terakhir kali sebelum program selesai dan juga dijalankan walaupun terjadi error di program.
	// pool.Close() adalah fungsi dari modul pgxpool yang digunakan untuk menutup koneksi ke database PostgreSQL.
	// jadi artinya walaupun terjadi error di program, koneksi ke database akan ditutup dengan normal dan tidak akan terjadi memory leak atau kebocoran memori.
	fmt.Println("✅ Sukses terhubung ke Database!")

	queries := db.New(pool)
	// queries ini digunakan untuk menampung remote koneksi ke database PostgreSQL.
	// db.New(pool) adalah fungsi dari modul db yang digunakan untuk membuat koneksi ke database PostgreSQL.
	// fungsi ini digunakan untuk mempermudah dalam melakukan query ke database PostgreSQL.
	_ = queries
	// _ = queries ini digunakan untuk menghilangkan warning dari golang, karena queries belum digunakan di program ini.

	r := gin.Default()
	// fungsi ini digunakan untuk membuat router baru dengan middleware default (logger dan recovery).
	// gin.Default() adalah fungsi dari modul gin yang digunakan untuk membuat router baru dengan middleware default (logger dan recovery).
	// r ini digunakan untuk menampung router baru yang dibuat oleh gin.Default().

	r.GET("/ping", func(c *gin.Context) {
	// r.GET() adalah fungsi dari modul gin untuk membuat route GET baru dengan path "/ping".
	// func(c *gin.Context) adalah fungsi handler yang bertugas untuk menangani request yang masuk ke route "/ping".
	// jadi artinya ketika ada request GET ke route "/ping", maka akan menjalankan fungsi handler di dalamnya.
		c.JSON(200, gin.H{
			"message": "pong! Server Perpustakaan Menyala 🚀",
		})
		// c.JSON() ini berfungsi untuk mengirimkan response JSON ke client dengan status code 200/success.
		// gin.H cara praktis untuk menyusun data JSON yang akan dikirim ke client.
	})

	fmt.Println("🚀 Server berjalan di http://localhost:8001")
	// fmt.Println() ini berfungsi untuk menampilkan log di console, tapi tidak menghentikan program.
	if err := r.Run(":8001"); err != nil {
	// r.Run(":8001") berfungsi untuk menjalankan server di port 8001
	// disini penggunaan if menggunakan (;) untuk pemisah:
	// sebelah kiri (;) digunakan untuk menjalankan program, sekaligus membuat variabel err untuk menampung error jika terjadi error saat menjalankan r.Run(":8080")
	// sebelah kanan (;) digunakan untuk mengecek apakah variabel err tidak nil/null/kosong, jika iya makan akan menjalankan {...}
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}


/*
Info Penting:
# Tanda (:=)
Tanda ini digunakan untuk mendeklarasikan variabel dan otomatis menganalisis tipe data dari isinya. Misal:
x := 10 // x otomatis bertipe int
y := "Hello" // y otomatis bertipe string

(:=) bisa digabungkan denga variebel baru dan variabel lama, misal:
x := 10 // x otomatis bertipe int
y := "Hello" // y otomatis bertipe string
x, z := 20, "World" // x otomatis bertipe int, z otomatis bertipe string

# defer
defer adalah fungsi yang dijalankan terakhir sebelum program selesai dengan normal, walaupun terjadi error didalam program. contoh:
---contoh code program berhasil---
func main() {
	fmt.Println("Buka Pintu")
	defer fmt.Println("Tutup Pintu") // ini akan dijalankan terakhir sebelum program selesai
	fmt.Println("Masuk Rumah")
}
---output---
Buka Pintu
Masuk Rumah
Tutup Pintu

---contoh code program error---
func main() {
	fmt.Println("Buka Pintu")
	defer fmt.Println("Tutup Pintu") // ini akan dijalankan terakhir sebelum program selesai
	panic("Terjadi Error") // ini akan menghentikan program secara paksa, tapi defer tetap dijalankan
	fmt.Println("Masuk Rumah") // program ini tidak pernah dijalankan karena program sudah dihentikan secara paksa oleh panic
}
---output---
Buka Pintu
Terjadi Error
Tutup Pintu

# Tanda (=)
Tanda ini digunakan untuk mengubah nilai dari variabel yang sudah dideklarasikan sebelumnya. Misal:
x = 20 // mengubah nilai x menjadi 20
y = "World" // mengubah nilai y menjadi "World"
*/