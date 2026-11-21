package main

import "fmt"

const NMAX int = 2002

type tabNIM [NMAX]int64

func cariNIM(T tabNIM, n int, x int64, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x && *idx == -1 {
			*idx = i
		}
	}
}
func ubahNIM(T *tabNIM, idx int, baru int64) {
	T[idx] = baru
}
func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var T tabNIM
	var n, idx int
	var cari, penggantiNIM int64

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&cari)
	cariNIM(T, n, cari, &idx)

	if idx != -1 {
		fmt.Scan(&penggantiNIM)
		ubahNIM(&T, idx, penggantiNIM)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx+1)
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}