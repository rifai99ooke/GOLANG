package main

import (
	"fmt"
	"math"
)

func luasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func luasPersegiPanjang(panjang,lebar float64) float64 {
	return panjang * lebar
}

func luasLingkaran (jariJari float64) float64 {
		return math.Pi * jariJari * jariJari
}
func LuasSegitiga(alas, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}
func LuasJajarGenjang(alas, tinggi float64) float64 {
	return alas * tinggi
}
func LuasBelahKetupat(diagonal1, diagonal2 float64) float64 {
	return 0.5 * diagonal1 * diagonal2
}
func luasLayangLayang(diagonal1, diagonal2 float64) float64 {
	return 0.5 * diagonal1 * diagonal2
}
func LuasTrapesium(alasAtas, alasBawah, tinggi float64) float64 {
	return 0.5 * (alasAtas + alasBawah) * tinggi
}
func LuasSegiLima(sisi, apotema float64) float64 {
	return 0.5 * sisi * apotema
}
func LuasSegiEnam(sisi float64) float64 {
	return (3 * math.Sqrt(3) / 2) * sisi * sisi
}

func main() {

	var pilihan int
	
	fmt.Println("Pilih opsi:")
	
	fmt.Println("1. Luas Persegi")
	
	fmt.Println("2. Luas Persegi Panjang")
	
	fmt.Println("3. Luas Lingkaran")
	
	fmt.Println("4. Luas Segitiga")
	
	fmt.Println("S. Luas Jajar Genjang")
	
	fmt.Println("6. Luas Belah Ketupat")
	
	fmt.Println("7. Luas Layang-Layang")
	
	fmt.Println("8. Luas Trapesium")
	
	fmt.Println("9. Luas Segi Lima")
	
	fmt.Println("10. Luas Segi Enam")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scanln(&pilihan)


	

switch pilihan {
case 1:
    var sisi float64
    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scanln(&sisi)
    luas := luasPersegi(sisi)
    fmt.Printf("Luas persegi: %.2f\n", luas)
case 2:
    var panjang, lebar float64
    fmt.Print("Masukkan panjang: ")
    fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)
	fmt.Printf("Luas Persegi Panjang: %.2f\n", luasPersegiPanjang(panjang, lebar))
case 3:
	var jariJari float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)
	fmt.Printf("Luas lingkaran: %.2f\n", luasLingkaran(jariJari))
case 4:
	var alas, tinggi float64
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)
	fmt.Printf("Luas segitiga: %.2f\n", LuasSegitiga(alas, tinggi))
case 5:
	var alas, tinggi float64
	fmt.Print("Masukkan alas jajar genjang: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi jajar genjang: ")
	fmt.Scan(&tinggi)
	fmt.Printf("Luas jajar genjang: %.2f\n", LuasJajarGenjang(alas, tinggi))
case 6:
	var diagonal1, diagonal2 float64
	fmt.Print("Masukkan diagonal 1 belah ketupat: ")
	fmt.Scan(&diagonal1)
	fmt.Print("Masukkan diagonal 2 belah ketupat: ")
	fmt.Scan(&diagonal2)
	fmt.Printf("Luas belah ketupat: %.2f\n", LuasBelahKetupat(diagonal1, diagonal2))
case 7:
	var diagonal1, diagonal2 float64
	fmt.Print("Masukkan diagonal 1 layang-layang: ")
	fmt.Scan(&diagonal1)
	fmt.Print("Masukkan diagonal 2 layang-layang: ")
	fmt.Scan(&diagonal2)
	fmt.Printf("Luas layang-layang: %.2f\n", luasLayangLayang(diagonal1, diagonal2))
case 8:
	var alasAtas, alasBawah, tinggi float64
	fmt.Print("Masukkan alas atas trapesium: ")
	fmt.Scan(&alasAtas)
	fmt.Print("Masukkan alas bawah trapesium: ")	
	fmt.Scan(&alasBawah)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)
	fmt.Printf("Luas trapesium: %.2f\n", LuasTrapesium(alasAtas, alasBawah, tinggi))
case 9:
	var sisi, apotema float64
	fmt.Print("Masukkan sisi segi lima: ")
	fmt.Scan(&sisi)
	fmt.Print("Masukkan apotema segi lima: ")
	fmt.Scan(&apotema)
	fmt.Printf("Luas segi lima: %.2f\n", LuasSegiLima(sisi, apotema))
case 10:
	var sisi float64
	fmt.Print("Masukkan sisi segi enam: ")
	fmt.Scan(&sisi)
	fmt.Printf("Luas segi enam: %.2f\n", LuasSegiEnam(sisi))
}
}