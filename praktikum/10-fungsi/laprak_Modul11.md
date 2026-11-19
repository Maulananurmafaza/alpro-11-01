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
 
 ## Unguided

### 1. nim.go

```go

```

##### Output



#### Deskripsi
Sistem akademik universitas menyimpan daftar NIM mahasiswa.lalu kita membuat program untuk mencari apakah seorang mahasiswa terdaftar berdasarkan NIM-nya. Jika mahasiswa ditemukan, NIM mahasiswa tersebut diubah dengan NIM pengganti

## Kesimpulan


## Referensi
