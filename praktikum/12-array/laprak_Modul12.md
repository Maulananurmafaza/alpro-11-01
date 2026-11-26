# <h1 align="center">Laporan Praktikum Modul 12 - array </h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Array


## Guided

### 1. curah.go

```go
package main

import "fmt"

const NMAX int = 9999

type arrReal [NMAX]float64

// Procedure bacaData(in/out A : arrReal, in n : integer)
func bacaData(A *arrReal, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

// Procedure cetakData(in A : arrReal, in n : integer)
func cetakData(A arrReal, n int) {
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%.1f.\n", A[i])
		} else {
			fmt.Printf("%.1f, ", A[i])
		}
	}
}

// Procedure cariMaksMin(in A : arrReal, in n : integer, out iMaks, iMin : integer)
func cariMaksMin(A arrReal, n int, iMaks, iMin *int) {
	*iMaks = 0
	*iMin = 0
	for i := 1; i < n; i++ {
		// Menggunakan > dan < (bukan >= atau <=) agar jika ada nilai sama, indeks pertama yang diambil
		if A[i] > A[*iMaks] {
			*iMaks = i
		}
		if A[i] < A[*iMin] {
			*iMin = i
		}
	}
}

// Function rataRata(A : arrReal, n : integer) -> real
func rataRata(A arrReal, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	return total / float64(n)
}

// Procedure balikData(in/out A : arrReal, in n : integer)
func balikData(A *arrReal, n int) {
	var temp float64
	for i := 0; i < n/2; i++ {
		temp = A[i]
		A[i] = A[n-1-i]
		A[n-1-i] = temp
	}
}

func main() {
	var data arrReal
	var n, iMaks, iMin int

	// Baca banyaknya data N
	fmt.Scan(&n)

	// Validasi batas N sesuai soal (1 <= N <= 9999)
	if n >= 1 && n <= NMAX {
		// 1. Baca N data curah hujan
		bacaData(&data, n)

		// 2. Cetak data awal
		fmt.Print("Data awal: ")
		cetakData(data, n)

		// 3. Cari dan cetak nilai tertinggi & terendah beserta posisinya (indeks + 1)
		cariMaksMin(data, n, &iMaks, &iMin)
		fmt.Printf("Tertinggi: %.1f (data ke-%d)\n", data[iMaks], iMaks+1)
		fmt.Printf("Terendah: %.1f (data ke-%d)\n", data[iMin], iMin+1)

		// 4. Hitung dan cetak rata-rata (2 angka di belakang koma)
		fmt.Printf("Rata-rata: %.2f\n", rataRata(data, n))

		// 5. Balik urutan elemen tanpa array bantuan dan cetak hasilnya
		balikData(&data, n)
		fmt.Print("Data setelah dibalik: ")
		cetakData(data, n)
	}
}
```
#### Deskripsi
 mengolah data curah hujan bulanan menggunakan array. Program membaca sejumlah data curah hujan, menampilkan data awal, mencari nilai tertinggi dan terendah beserta posisinya, menghitung rata-rata curah hujan, serta membalik urutan data tanpa menggunakan array bantuan.

 ## Unguided

### 1. perpustakaan.go

```go
package main

import "fmt"

const NMAX int = 100
type book struct {
	code string
	title string
	borrowed int
	pages int
}

type books [NMAX]book

func bacaData (T *books, n *int) {
	*n = 0
	fmt.Scan(&T[*n].code)

	for T[*n].code != "SELESAI" {
		fmt.Scan(&T[*n].title,
			&T[*n].borrowed,
			&T[*n].pages)
		*n++
		fmt.Scan(&T[*n].code)
	}
}

func cetakData (T books, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].code,
			T[i].title,
			T[i].borrowed,
			T[i].pages)
	}
}

func hapusBuku (T *books, n *int, X int) {
	i := 0

	for i < *n {
		if T[i].borrowed < X {
			for j := i; j < *n-1; j++ {
				T[j] = T[j+1]
			}
			*n--
		} else {
			i++
		}
	}
}

func totalPinjam (T books, n int) int {
	total := 0

	for i := 0; i < n; i++ {
		total += T[i].borrowed
	}

	return total
}

func idxFavorit(T books, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if T[i].borrowed > T[idx].borrowed ||
			(T[i].borrowed == T[idx].borrowed &&
				T[i].pages < T[idx].pages) {
			idx = i
		}
	}

	return idx
}

func main() {
    var T books
    var n, X int

    bacaData(&T, &n)

    fmt.Println("Data awal buku:")
    cetakData(T, n)

    fmt.Scan(&X)

    hapusBuku(&T, &n, X)

    fmt.Println("Data setelah penghapusan:")

    if n == 0 {
        fmt.Println("Tidak ada buku tersisa")
        return
    }

    cetakData(T, n)

    fmt.Println("Total peminjaman:",
        totalPinjam(T, n))

    idx := idxFavorit(T, n)

    fmt.Println("Buku paling sering dipinjam:")
	fmt.Println(T[idx].code,
    T[idx].title,
    T[idx].borrowed,
    T[idx].pages)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/12-array/unguided/perpustakaan_output.png

#### Deskripsi
Sebuah perpustakaan ingin merapikan data bukunya. Data setiap buku terdiri dari kode buku, judul buku, jumlah peminjaman, dan jumlah halaman. Petugas ingin menghapus dari array semua buku yang jumlah peminjamannya kurang dari batas X, lalu menampilkan total peminjaman dan buku paling sering dipinjam dari data yang tersisa. Jika terdapat beberapa buku dengan jumlah peminjaman tertinggi yang sama, pilih buku dengan jumlah halaman lebih sedikit.
## Kesimpulan
array dapat digunakan untuk menyimpan dan mengolah sekumpulan data secara terstruktur. Pada soal curah hujan, array digunakan untuk mencari nilai tertinggi dan terendah, menghitung rata-rata, serta membalik urutan data. Pada soal perpustakaan, array digunakan untuk menghapus data buku berdasarkan batas peminjaman, menghitung total peminjaman, dan menentukan buku yang paling sering dipinjam. Kedua soal juga melatih penggunaan subprogram berupa prosedur dan fungsi untuk membuat program lebih terstruktur dan mudah dipahami.
## Referensi
1. Google. (n.d.). A Tour of Go: Arrays. https://go.dev/tour/moretypes/6
2. Google. (n.d.). Go by Example: Arrays. https://gobyexample.com/arrays
3. Google. (n.d.). Go by Example: Slices. https://gobyexample.com/slices