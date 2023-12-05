package main

import (
	"fmt"
	"math"
)

func main() {
	var jam1, mint1, jam2, mint2, harga, tahun, tanggal, jam, tanggal1, tahun1, tanggal2, tahun2, waktu, bulan, bulan1, bulan2 float64
	var bul2, bul1 string
	fmt.Println("Waktu penitipan")
	fmt.Println("TANGGAL")
	fmt.Scanln(&tanggal1, &bul1, &tahun1)
	fmt.Println("JAM")
	fmt.Scanln(&jam1, &mint1)

	fmt.Println("Waktu pengambilan")
	fmt.Println("TANGGAL")
	fmt.Scanln(&tanggal2, &bul2, &tahun2)
	fmt.Println("JAM")
	fmt.Scanln(&jam2, &mint2)

	if bul1 == "Januari" || bul2 == "Januari" {
		bulan1 = float64(1 * 31 * 24)
		bulan2 = float64(1 * 31 * 24)
	} else if bul1 == "Februari" {
		bulan1 = float64(2 * 30 * 24)
		bulan2 = float64(2 * 30 * 24)
	} else if bul1 == "Maret" {
		bulan1 = float64(3 * 31 * 24)
		bulan2 = float64(3 * 31 * 24)
	} else if bul1 == "April" {
		bulan1 = float64(4 * 30 * 24)
		bulan2 = float64(4 * 30 * 24)
	} else if bul1 == "Mei" {
		bulan1 = float64(5 * 31 * 24)
		bulan2 = float64(5 * 31 * 24)
	} else if bul1 == "Juni" {
		bulan1 = float64(6 * 30 * 24)
		bulan2 = float64(6 * 30 * 24)
	} else if bul1 == "Juli" {
		bulan1 = float64(7 * 31 * 24)
		bulan2 = float64(7 * 31 * 24)
	} else if bul1 == "Agustus" {
		bulan1 = float64(8 * 31 * 24)
		bulan2 = float64(8 * 31 * 24)
	} else if bul1 == "September" {
		bulan1 = float64(9 * 30 * 24)
		bulan2 = float64(9 * 30 * 24)
	} else if bul1 == "Oktober" {
		bulan1 = float64(10 * 31 * 24)
		bulan2 = float64(10 * 31 * 24)
	} else if bul1 == "November" {
		bulan1 = float64(11 * 30 * 24)
		bulan2 = float64(11 * 30 * 24)
	} else if bul1 == "Desember" {
		bulan1 = float64(12 * 31 * 24)
		bulan2 = float64(12 * 31 * 24)
	}

	if bul2 == "Januari" {
		bulan2 = float64(1 * 31 * 24)
	} else if bul2 == "Februari" {
		bulan2 = float64(2 * 30 * 24)
	} else if bul2 == "Maret" {
		bulan2 = float64(3 * 31 * 24)
	} else if bul2 == "April" {
		bulan2 = float64(4 * 30 * 24)
	} else if bul2 == "Mei" {
		bulan2 = float64(5 * 31 * 24)
	} else if bul2 == "Juni" {
		bulan2 = float64(6 * 30 * 24)
	} else if bul2 == "Juli" {
		bulan2 = float64(7 * 31 * 24)
	} else if bul2 == "Agustus" {
		bulan2 = float64(8 * 31 * 24)
	} else if bul2 == "September" {
		bulan2 = float64(9 * 30 * 24)
	} else if bul2 == "Oktober" {
		bulan2 = float64(10 * 31 * 24)
	} else if bul2 == "November" {
		bulan2 = float64(11 * 30 * 24)
	} else if bul2 == "Desember" {
		bulan2 = float64(12 * 31 * 24)
	}

	bulan = bulan2 - bulan1
	tahun = (tahun2 - tahun1) * 12 * 30 * 24
	tanggal = (tanggal2 - tanggal1) * 24
	jam = math.Abs((jam2*60+mint2)-(jam1*60+mint1)) / 60
	waktu = tahun + bulan + tanggal + jam
	harga = waktu * 500

	fmt.Println("RP.", harga)
}
