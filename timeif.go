package main

import (
	"fmt"
	"math"
)

func main() {
	var jam1, mint1, jam2, mint2, harga, tahun, bulan, tanggal, jam, tanggal1, bulan1, tahun1, tanggal2, bulan2, tahun2, waktu float64
	fmt.Println("Waktu penitipan")
	fmt.Println("TANGGAL")
	fmt.Scanln(&tanggal1, &bulan1, &tahun1)
	fmt.Println("JAM")
	fmt.Scanln(&jam1, &mint1)

	fmt.Println("Waktu pengambilan")
	fmt.Println("TANGGAL")
	fmt.Scanln(&tanggal2, &bulan2, &tahun2)
	fmt.Println("JAM")
	fmt.Scanln(&jam2, &mint2)
	//Tahun
	if tahun2 >= tahun1 {
		tahun = (tahun2 - tahun1) * 12 * 30 * 24
	} else {
		fmt.Println("Ketentuan Tahun Salah")
	}

	//Bulan
	if (01 <= bulan1 && bulan1 <= 12) && (01 <= bulan2 && bulan2 <= 12) {
		bulan1 = bulan1 * 30 * 24
		bulan2 = bulan2 * 30 * 24
		bulan = bulan2 - bulan1
	} else {
		fmt.Println("Ketentuan Bulan salah")
	}

	//Tanggal
	if (01 <= tanggal1 && tanggal1 <= 30) && (01 <= tanggal2 && tanggal2 <= 30) {
		tanggal1 = tanggal1 * 24
		tanggal2 = tanggal2 * 24
		tanggal = tanggal2 - tanggal1
	} else {
		fmt.Println("Ketentuan Tanggal salah")
	}

	//Jam
	if (00 <= jam1 && jam1 <= 24) && (00 <= jam2 && jam2 <= 24) {
		jam1 = jam1*60 + mint1
		jam2 = jam2*60 + mint2
		jam = math.Abs(jam2-jam1) / 60
	} else {
		fmt.Println("Ketentuan Jam salah")
	}

	// Total
	if tahun == 0 && bulan == 0 && tanggal == 0 && jam == 0 {
		fmt.Println("Coba lakukan lagi")
	} else {
		waktu = tahun + bulan + tanggal + jam
		harga = waktu * 500
		fmt.Println("RP.", harga)
	}
}
