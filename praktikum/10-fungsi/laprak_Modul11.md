# <h1 align="center">Laporan Praktikum Modul 11 - Prosedur</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Prosedur
Prosedur dapat dianggap sebagai potongan beberapa instruksi program menjadi suatu instruksi baru yang dibuat untuk mengurangi kerumitan dari kode program yang kompleks pada suatu program yang besar. Prosedur akan menghasilkan suatu akibat atau efek langsung pada program ketika dipanggil pada program utama.

## Guided

### 1. sapi.go

```go
package main

import "fmt"

const NMAX = 10000

type Arr [NMAX]int

func mengisiSapi(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func geserKanan(T *Arr, n int) {
	var temp int = T[n-1]

	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}

	T[0] = temp
}

func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var T Arr
	mengisiSapi(&T, n)

	k = k % n

	for i := 0; i < k; i++ {
		geserKanan(&T, n)
	}

	cetak(T, n)

}
```
#### Deskripsi
 Menentukan susunan nomor sapi pada (a) kandang setelah digeser ke kanan setiap hari selama (b) hari, dengan memanfaatkan modulo a % b agar tidak mengulang pergeseran berlebih.

 ## Unguided

### 2. nim.go

```go
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
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/10-fungsi/unguided_2/nim_output.png

#### Deskripsi
Sistem akademik universitas menyimpan daftar NIM mahasiswa.lalu kita membuat program untuk mencari apakah seorang mahasiswa terdaftar berdasarkan NIM-nya. Jika mahasiswa ditemukan, NIM mahasiswa tersebut diubah dengan NIM pengganti

## Kesimpulan
Prosedur di Go digunakan sebagai func tanpa tipe nilai balik yang memisahkan alur data lewat mekanisme passing parameter: nilai biasa (by value) untuk mode in dan pointer untuk mode out atau in/out agar perubahan pada variabel, array, maupun struct terbaca di pemanggil.

## Referensi
1. Dokumentasi Resmi Bahasa Go: The Go Authors. The Go Programming Language Specification — bagian Function declarations, Pointer types, dan Array types. Tersedia di: [https://go.dev/ref/spec](https://go.dev/ref/spec)