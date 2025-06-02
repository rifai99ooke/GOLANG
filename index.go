package main

import (
	"fmt"
)
 func main () {
var totalbelanja float64
fmt.Println("masukkan total belanja anda = ")
fmt.Scan(&totalbelanja)

// variable data
var diskon float64
var totalbayar float64

// menentukan besaran diskon
if totalbelanja > 500000 {
	diskon = totalbelanja * 0.20
	fmt.Println("selamat anda mendapatkan diskon 20%")
} else if totalbelanja >= 200000 && totalbelanja <= 500000 {
	diskon = totalbelanja * 0.10
	fmt.Println("selamat anda mendapatkan diskon 10%")
}
// menghitung total yang harus dibayar setelah diskon
totalbayar = totalbelanja - diskon
fmt.Printf("total belanja = Rp%.2f\n", totalbelanja)
fmt.Printf("diskon = Rp%.2f\n", diskon)
fmt.Printf("total bayar = Rp%.2f\n", totalbayar)
 }