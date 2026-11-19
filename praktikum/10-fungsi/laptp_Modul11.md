# <h1 align="center">Tugas Pendahuluan Modul 11 - Prosedur </h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. sekolah.go

```go
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

```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/11-prosedur/tp/sekolah/sekolah_output.png

#### Deskripsi
Sebuah sekolah menyimpan daftar nama murid berdasarkan urutan absen. lalu membuat program untuk mencari apakah seorang murid terdaftar, sekaligus menghitung banyaknya murid yang namanya diawali huruf yang sama dengan huruf pertama nama yang dicari.

### 2. penduduk.go

```go
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

```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/11-prosedur/tp/penduduk/penduduk_output.png

#### Deskripsi
Program ini digunakan untuk menyimpan data penduduk yang terdiri dari nama, tahun lahir, dan kota asal. Program akan mencari data penduduk berdasarkan tahun lahir, lalu menampilkan penduduk tertua dan termuda.

## Kesimpulan
prosedur dalam bahasa Go digunakan untuk membuat program lebih terstruktur dan mudah dipahami.kita belajar menggunakan prosedur untuk mencari data, menghitung jumlah, serta menentukan data tertua dan termuda