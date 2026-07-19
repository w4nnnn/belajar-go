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
	if err != nil {
	// variabel err tidak nil/null/kosong maka akan menjalankan dibawah ini.
	// jadi artinya jika output dari godotenv.Load() null atau kosong, program berhasil berjalan. tapi jika ada output/isinya berarti program ada yang error
		log.Fatalf("Gagal load file .env: %v\n", err)
		// log menampilkan/print log error.
		// Fatalf berfungsi mencetak/print sesuai format yang ditentukan, lalu menghentikan program.
		// %v ini adala tempat kosong yang akan di gantikan dengan nilai yang akan diberikan setelah koma, dalam hal ini adalah err.
	}

	dbURL := os.Getenv("GOOSE_DBSTRING")
	// dbURL adalah variabel yang akan menampung nilai dari os.Getenv("GOOSE_DBSTRING"), yaitu mengambil nilai dari file .env dengan key GOOSE_DBSTRING.
	if dbURL == "" {
	// jika dbURL kosong, maka akan menjalankan dibawah ini.
		log.Fatal("GOOSE_DBSTRING tidak ditemukan di file .env")
		// log.Fatal akan menampilkan log error dan menghentikan program.
		// disini tidak menggunakan Fatalf karena tidak ada format tertentu yang akan ditampilkan yang cukup string saja.
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Gagal connect ke database: %v\n", err)
	}
	defer pool.Close()
	fmt.Println("✅ Sukses terhubung ke Database!")

	queries := db.New(pool)
	_ = queries

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong! Server Perpustakaan Menyala 🚀",
		})
	})

	fmt.Println("🚀 Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
