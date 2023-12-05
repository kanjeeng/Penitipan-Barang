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

	tahun = (tahun2 - tahun1) * 12 * 30 * 24
	bulan = (bulan2 - bulan1) * 30 * 24
	tanggal = (tanggal2 - tanggal1) * 24
	jam = math.Abs((jam2*60+mint2)-(jam1*60+mint1)) / 60

	waktu = tahun + bulan + tanggal + jam
	harga = waktu * 500

	fmt.Println("RP.", harga)
}
