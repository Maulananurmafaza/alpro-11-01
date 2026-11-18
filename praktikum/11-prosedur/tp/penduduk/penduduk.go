package main

import "fmt"

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

func cariTahun(T []penduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			return
		}
	}
}
func tertuaTermuda(T []penduduk, n int, tua, muda *int) {
	*tua = 0
	*muda = 0

	for i := 1; i < n; i++ {
		if T[i].tahun < T[*tua].tahun {
			*tua = i
		}
		if T[i].tahun > T[*muda].tahun {
			*muda = i
		}
	}
}
func main() {
	var n, th int
	var T [1000]penduduk

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}
	fmt.Scan(&th)

	idx, tua, muda := 0, 0, 0

	cariTahun(T[:], n, th, &idx)
	tertuaTermuda(T[:], n, &tua, &muda)

	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di indeks ke-%d\n",
			T[idx].nama, T[idx].tahun, T[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	fmt.Printf("Penduduk tertua: %s (%d)\n",
		T[tua].nama, T[tua].tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n",
		T[muda].nama, T[muda].tahun)
}
