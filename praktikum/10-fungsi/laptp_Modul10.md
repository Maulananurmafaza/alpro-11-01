# <h1 align="center">Tugas Pendahuluan Modul 10 - Fungsi </h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. raport.go

```go
package main

import "fmt"

type rapot struct {
	nama string
	nilai int
}

func nilaiHuruf(x int) string {
	if x > 80 {
		return "A"
	} else if x > 70 {
		return "AB"
	} else if x > 65 {
		return "B"
	} else if x > 60 {
		return "BC"
	} else if x > 50 {
		return "C"
	} else if x > 40 {
		return "D"
	}
	return "E"
}

func rataRata(T []rapot, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}

func main() {
	var n int
	fmt.Scan(&n)

	var T [999]rapot

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].nilai)
	}
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nama, T[i].nilai, nilaiHuruf(T[i].nilai))
	}

	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T[:], n))
}
```

##### Output


#### Deskripsi
Seorang mahasiswa memiliki data raport yang berisi daftar mata kuliah dan nilai angka yang diperoleh lalu mebuat program yang menampilkan setiap mata kuliah beserta nilai huruf-nya, kemudian menampilkan rata-rata nilai seluruh mata kuliah. Konversi nilai huruf dan perhitungan rata-rata

### 2. inputNilai.go

```go
package main

import "fmt"

type student struct {
	nim int
	nama string
	grade float64
}

func cariNIM(T []student, n, nim int) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}
func rataRata(T []student, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}
func main() {
	var T [100]student
	n := 0

	for {
		var input, nama string
		var nilai float64
		var nim int

		fmt.Scan(&input)

		if input == "STOP" {
			break
		}

		fmt.Sscan(input, &nim)
		fmt.Scan(&nama, &nilai)

		if cariNIM(T[:], n, nim) == -1 {
			T[n] = student{nim, nama, nilai}
			n++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].nama, T[i].grade)
	}
	if n > 0 {
		fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T[:], n))
	}
}
```

##### Output

#### Deskripsi
Soal ini meminta kita membuat program untuk memasukkan data mahasiswa berupa NIM, nama, dan nilai. Program harus memeriksa NIM agar tidak ada data yang duplikat. Jika NIM sudah pernah dimasukkan, data tidak disimpan lagi. Input berhenti ketika memasukkan STOP, kemudian program menampilkan data mahasiswa dan menghitung rata-rata nilainya.

## Kesimpulan
fungsi dan prosedur dalam bahasa Go membantu mempermudah pembuatan program agar lebih terstruktur dan mudah dipahami.Kita belajar menggunakan fungsi untuk mengolah data, menghitung rata-rata, mencari data, serta menghindari duplikasi.
