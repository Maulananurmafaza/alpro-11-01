package main

import "fmt"

func cariMurid(T []string, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false

	for i := 0; i < n; i++ {
		if T[i] == nama {
			*posisi = i + 1
			*ketemu = true
			return
		}
	}
}
func hitungHuruf(T []string, n int, h byte, jumlah *int) {
	*jumlah = 0

	for i := 0; i < n; i++ {
		if T[i][0] == h {
			*jumlah++
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)

	T := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var nama string
	fmt.Scan(&nama)

	posisi := 0
	ketemu := false
	jumlah := 0

	cariMurid(T, n, nama, &posisi, &ketemu)
	hitungHuruf(T, n, nama[0], &jumlah)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", nama[0], jumlah)
}
