package main

import "fmt"

func main() {
	var jenis, durasi int
	var motor, mobil, total int
	var biaya int

	fmt.Println(" tarif parkir ")

	for {
		fmt.Println("1. Motor")
		fmt.Println("2. Mobil")
		fmt.Println("0. Selesai")
		fmt.Print("Pilih kendaraan: ")
		fmt.Scan(&jenis)

		if jenis == 0 {
			break
		}
		if jenis != 1 && jenis != 2 {
			fmt.Println("Jenis kendaraan tidak valid!")
			continue
		}
		fmt.Print("Masukkan durasi parkir (jam): ")
		fmt.Scan(&durasi)

		if durasi <= 0 {
			fmt.Println("Durasi tidak valid!")
			continue
		}
		if jenis == 1 {
			biaya = 2000 + (durasi-1)*1000
			if biaya > 10000 {
				biaya = 10000
			}
			motor++
			fmt.Println("Jenis: Motor")
		} else {
			biaya = 5000 + (durasi-1)*3000
			if biaya > 25000 {
				biaya = 25000
			}
			mobil++
			fmt.Println("Jenis: Mobil")
		}

		total += biaya

		fmt.Println("Biaya parkir: Rp", biaya)
	}

	fmt.Println("=== REKAP PARKIR ===")
	fmt.Println("Jumlah motor:", motor)
	fmt.Println("Jumlah mobil:", mobil)
	fmt.Println("Total pendapatan: Rp", total)
}