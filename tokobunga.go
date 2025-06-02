package main

import (
	"fmt"
)
func main() {
	

type jenisbunga struct {
	nama string
	harga int
}

listbunga := []jenisbunga{
	{"bunga matahari", 50000},
	{"bunga melati", 100000},
	{"bunga merah", 200000},
}
fmt.Println("Daftar harga bunga per tangkai:")
for i, bunga := range listbunga {
	fmt.Printf("%d. %s - %d\n", i+1, bunga.nama, bunga.harga)
}
var pilihbunga []int
fmt.Println("Masukkan nomor bunga yang ingin dibeli (ketik '0' untuk selesai):")
for {
	var belanjabunga int
	fmt.Print("Masukkan nomor bunga: ")
	fmt.Scan(&belanjabunga)
	if belanjabunga == 0 {
		break
	}
	pilihbunga = append(pilihbunga, belanjabunga)
}
var tangkai int
fmt.Print("Masukkan jumlah tangkai: ")
fmt.Scan(&tangkai)

totalbayar := 0
for _, belanjabunga := range pilihbunga {
	totalbayar += listbunga[belanjabunga-1].harga *tangkai
}
fmt.Printf("Total bayar: %d\n", totalbayar)

}