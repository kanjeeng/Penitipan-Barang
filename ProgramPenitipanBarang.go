package main

import (
	"fmt"
	"math"
)

func main() {

	// Variabel
	var totalbayar int
	var jumlahbarang, i float64
	var nama, barang, kondisi string
	var valid, tahun_valid, bulan_valid, tanggal_valid, jam_valid bool
	var hour1, hour2, date1, date2, month1, month2 float64
	var jam1, menit1, jam2, menit2, tahun, bulan, tanggal, jam, tanggal1, bulan1, tahun1, tanggal2, bulan2, tahun2, waktu, totalberat, beratbarang, bayarberat, bayarwaktu float64

	valid = false

	for valid != true {

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

		// Waktu Penitipan
		fmt.Println(" ")
		fmt.Println("Waktu penitipan")
		fmt.Println("*Contoh Penulisan Tanggal = 'tanggal <spasi> bulan <spasi> tahun'")
		fmt.Println("*Contoh Penulisan Jam = 'jam <spasi> menit'")
		fmt.Println(" ")
		fmt.Print("Tanggal = ")
		fmt.Scan(&tanggal1, &bulan1, &tahun1)
		fmt.Print("Jam = ")
		fmt.Scan(&jam1, &menit1)

		fmt.Println(" ")
		fmt.Println("Waktu pengambilan")
		fmt.Println("*Contoh Penulisan Tanggal = 'tanggal <spasi> bulan <spasi> tahun'")
		fmt.Println("*Contoh Penulisan Jam = 'jam <spasi> menit'")
		fmt.Println(" ")
		fmt.Print("Tanggal = ")
		fmt.Scan(&tanggal2, &bulan2, &tahun2)
		fmt.Print("Jam = ")
		fmt.Scan(&jam2, &menit2)

		//Tahun
		if tahun2 >= tahun1 {
			tahun = (tahun2 - tahun1) * 12 * 30 * 24
			if tahun2-tahun1 >= 0 {
				tahun_valid = true
			} else {
				fmt.Println("Ketentuan Tahun Salah")
				tahun_valid = false
			}
		} else {
			fmt.Println("Ketentuan Tahun Salah")
			tahun_valid = false
		}

		//Bulan
		if (01 <= bulan1 && bulan1 <= 12) && (01 <= bulan2 && bulan2 <= 12) {
			month1 = bulan1 * 30 * 24
			month2 = bulan2 * 30 * 24
			bulan = month2 - month1
			if month2-month1 >= 0 {
				bulan_valid = true
			} else {
				fmt.Println("Ketentuan Bulan salah")
				bulan_valid = false
			}
		} else {
			fmt.Println("Ketentuan Bulan salah")
			bulan_valid = false
		}

		//Tanggal
		if (01 <= tanggal1 && tanggal1 <= 30) && (01 <= tanggal2 && tanggal2 <= 30) {
			date1 = tanggal1 * 24
			date2 = tanggal2 * 24
			tanggal = date2 - date1
			if date2-date1 >= 0 {
				tanggal_valid = true
			} else {
				fmt.Println("Ketentuan Tanggal salah")
				tanggal_valid = false
			}
		} else {
			fmt.Println("Ketentuan Tanggal salah")
			tanggal_valid = false
		}

		//Jam
		if (00 <= jam1 && jam1 <= 24) && (00 <= jam2 && jam2 <= 24) {
			hour1 = jam1*60 + menit1
			hour2 = jam2*60 + menit2
			jam = math.Abs(hour2-hour1) / 60
			if hour2-hour1 >= 0 {
				jam_valid = true
			} else {
				fmt.Println("Ketentuan Jam salah")
				jam_valid = false
			}
		} else {
			fmt.Println("Ketentuan Jam salah")
			jam_valid = false
		}

		// Total
		if tahun_valid && bulan_valid && tanggal_valid && jam_valid {
			waktu = tahun + bulan + tanggal + jam
			bayarwaktu = waktu * 500

			// Total Bayar
			bayarberat = totalberat * 1000
			totalbayar = int(bayarwaktu) + int(bayarberat)

			// Output Struk
			fmt.Println(" ")
			fmt.Println("--- Struk Pembayaran ---")
			fmt.Println(" ")

			// Output
			fmt.Println("Nama : ", nama)
			fmt.Println("Waktu Penitipan : ")
			fmt.Printf("      Tanggal = %g-%g-%g \n ", tanggal1, bulan1, tahun1)
			fmt.Printf("     Jam     = %g.%g \n", jam1, menit1)
			fmt.Println("Waktu Pengambilan : ")
			fmt.Printf("      Tanggal = %g-%g-%g \n ", tanggal2, bulan2, tahun2)
			fmt.Printf("     Jam     = %g.%g \n ", jam2, menit2)
			fmt.Println("Total Bayar : Rp.", totalbayar)
			fmt.Println(" ")

			fmt.Println("------------------------")

			fmt.Println(" ")
			fmt.Println("Kondisi Barang mengalami kerusakan atau tidak (iya/tidak)")
			fmt.Scan(&kondisi)

			if kondisi == "iya" {
				fmt.Println("Jika barang mengalami kerusakkan, \nmaka akan mendapatkan diskon sebesar 50%")
				diskon := float64(totalbayar) * 0.5
				fmt.Println("Total Pembayaran = Rp.", totalbayar-int(diskon))
			} else if kondisi == "tidak" {
				fmt.Println("Silahkan membayar Rp.", totalbayar)
			}

			valid = true
		} else {
			fmt.Println("Coba lakukan lagi")
			fmt.Println(" ")
			valid = false
		}

	}

}
