package main

import (
	"fmt"
	"time"
)

func main() {
	// Input data jenis kendaraan
	var jenisKendaraan string
	fmt.Print("Masukkan jenis kendaraan (Mobil/Motor): ")
	fmt.Scanln(&jenisKendaraan)

	// Menentukan harga parkir kendaraan
	var hargaParkir int
	if jenisKendaraan == "Mobil" {
		hargaParkir = 1500
	} else if jenisKendaraan == "Motor" {
		hargaParkir = 500
	} else {
		fmt.Println("Jenis kendaraan yang anda masukkan salah!")
		return
	}

	// Input data jam masuk dan keluar
	var jamMasuk, jamKeluar string
	fmt.Print("Masukkan jam masuk (hh:mm): ")
	fmt.Scanln(&jamMasuk)
	fmt.Print("Masukkan jam keluar (hh:mm): ")
	fmt.Scanln(&jamKeluar)

	// Parsing input waktu
	layout := "15:04"
	masuk, err := time.Parse(layout, jamMasuk)
	if err != nil {
		fmt.Println("Format jam masuk tidak valid!")
		return
	}
	keluar, err := time.Parse(layout, jamKeluar)
	if err != nil {
		fmt.Println("Format jam keluar tidak valid!")
		return
	}

	// Menghitung lama parkir dalam menit
	durasi := keluar.Sub(masuk).Minutes()
	if durasi < 0 {
		fmt.Println("Jam keluar tidak boleh lebih awal dari jam masuk!")
		return
	}

	// Menghitung biaya parkir
	biaya := hargaParkir * int(durasi) / 60

	// Menampilkan hasil
	fmt.Printf("Jenis kendaraan: %s\n", jenisKendaraan)
	fmt.Printf("Harga parkir per jam: %d\n", hargaParkir)
	fmt.Printf("Lama parkir: %.0f menit\n", durasi)
	fmt.Printf("Biaya parkir: %d\n", biaya)
}
