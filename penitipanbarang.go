package main

import (
	"fmt"
)

func main() {

	// Variabel
	var jumlahbarang, i, totalbayar, jenis int
	var nama, barang string
	var beratbarang, totalberat, bayarberat, jam1, menit1, jam2, menit2, tjam1, tjam2, hari, waktu, bayarwaktu float64

	// Definisi Variabel
	i = 1
	totalberat = 0

	// Nama Program
	fmt.Println(" ")
	fmt.Println("--- Program penitipan barang ---")
	fmt.Println(" ")

	// Masukkan Nama Penitip Barang
	fmt.Print("Nama : ")
	fmt.Scan(&nama)

	// Masukkan Jumlah Barang
	fmt.Print("Jumlah Barang yang akan dititipkan : ")
	fmt.Scan(&jumlahbarang)
	fmt.Println(" ")

	// Nama & Berat Barang
	for i <= jumlahbarang {
		fmt.Print("Nama Barang ", i, " : ")
		fmt.Scan(&barang)
		fmt.Print("Berat Barang (Kg) : ")
		fmt.Scan(&beratbarang)
		totalberat = totalberat + beratbarang
		i++
	}

	// Pilih Paket Waktu
	fmt.Println(" ")
	fmt.Println("--- Paket Waktu ---")
	fmt.Println("(1) Jam")
	fmt.Println("(2) Harian")
	fmt.Print("Pilih Jenis Waktu Penitipan (1/2) : ")
	fmt.Scan(&jenis)

	// If Else Paket
	// Paket Jam
	if jenis == 1 {
		fmt.Println("Anda memilih paket Jam")

		fmt.Println("Waktu penitipan")
		fmt.Println("Catatan : Contoh penulisan 'jam <spasi> menit' ")
		fmt.Scan(&jam1, &menit1)

		fmt.Println("Waktu pengambilan")
		fmt.Println("*Catatan : Contoh penulisan 'jam <spasi> menit' ")
		fmt.Scan(&jam2, &menit2)

		// Penitipan
		tjam1 = jam1*60 + menit1
		tjam2 = jam2*60 + menit2

		waktu = (tjam2 - tjam1) / 60

		// Output Struk
		fmt.Println(" ")
		fmt.Println("--- Struk Pembayaran ---")
		fmt.Println(" ")

		// Output Waktu
		fmt.Println("Waktu Penitipan : ", "jam", jam1, "menit", menit1)
		fmt.Println("Waktu Pengambilan : ", "jam", jam2, "menit", menit2)

		// Paket Harian
	} else if jenis == 2 {
		fmt.Println("Anda memilih paket Harian")

		fmt.Print("Berapa hari penitipan : ")
		fmt.Scan(&hari)

		// Penitipan
		waktu = hari * 24

		// Output Struk
		fmt.Println(" ")
		fmt.Println("--- Struk Pembayaran ---")
		fmt.Println(" ")

		// Output Waktu
		fmt.Println("Waktu Penitipan : ", hari, "hari")
	}

	// Total Bayar
	bayarberat = totalberat * 1000
	bayarwaktu = waktu * 500
	totalbayar = int(bayarwaktu) + int(bayarberat)

	// Output
	fmt.Println("Total Berat Barang : ", totalberat, "kg")
	fmt.Println("Total Bayar : Rp.", totalbayar)
	fmt.Println(" ")

	fmt.Println("------------------------")

}
