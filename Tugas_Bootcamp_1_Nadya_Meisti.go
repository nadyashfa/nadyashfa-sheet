package main

import "fmt"

func main() {
	var panjang = 3
	var lebar = 4
	var luaspersegi = panjang * lebar

	var alas = 4
	var tinggi = 2
	var luassegitiga = (alas * tinggi) / 2

	var jarijari = 7
	var phi = 22 / 7
	var luaslingkaran = jarijari * jarijari * phi

	var massa = 6
	var kecepatan = 2
	var energikinetik = 1/2*massa*kecepatan ^ 2
	var gravitasi = 10
	var ketinggian = 5
	var energipotensial = massa * gravitasi * ketinggian

	var jumlahgetaran = 15
	var waktu = 5
	var frekuensi = jumlahgetaran / waktu

	var celcius = 27
	var fahrenheit = (celcius * 9 / 5) + 32

	fmt.Println(luaspersegi)
	fmt.Println(luassegitiga)
	fmt.Println(luaslingkaran)
	fmt.Println(energikinetik)
	fmt.Println(energipotensial)
	fmt.Println(fahrenheit)
	fmt.Println(frekuensi)

}
