# <h1 align="center">Laporan Praktikum Modul 10 - fungsi</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Fungsi
Fungsi merupakan satu kesatuan rangkaian intruksi yang memberikan atau menghasilkan suatu nilai yang dimana biasanya memetakan input ke suatu nilai yang lain

## Guided

### 1. player.go

```go
package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

func idxTertinggi(T tabPlayer, n int) int {
	idx := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idx].score {
			idx = i
		}
	}
	return idx
}

func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var T tabPlayer
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	average := rataRata(T, n)
	best := idxTertinggi(T, n)
	di := hitungDiAtas(T, n, average)

	fmt.Printf("Rata-rata skor: %.2f\n", average)
	fmt.Printf("Player terbaik: %s (%d)\n", T[best].nickname, T[best].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", di)
}
```
#### Deskripsi
 program ini untuk mengolah data skor player sebuah turnamen game. Program harus menampilkan rata-rata skor, player dengan skor tertinggi, dan jumlah player yang skornya di atas rata-rata
 ## Unguided

### 1. gempa.go

```go
package main 

import "fmt"

const NMAX int = 999
type gempa struct {
	nama string
	tahun int
	magnitudo float64
}
type tabGempa [NMAX] gempa

func kategori (m float64)string {
	if m < 4.0 {
		return "Minor"
	} else if m >= 4.0 && m < 6.0 {
		return "Sedang"
	} else if m >= 6.0 && m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}
func rataMagnitudo(T tabGempa, n int) float64 {
    jumlah := 0.0

    for i := 0; i < n; i++ {
        jumlah += T[i].magnitudo
    }
    return jumlah / float64(n)
}
func idxTerkuat (T tabGempa, n int)int {
	idx := 0

	for i := 1; i < n; i++ {
        if T[i].magnitudo > T[idx].magnitudo {
            idx = i
        }
    }
    return idx
}
func hitungKategori (T tabGempa, n int, k string) int {
	jumlah := 0

	for i := 0; i < n; i++ {
        if kategori(T[i].magnitudo) == k {
            jumlah++
        }
    }
    return jumlah
}
func main(){
	var T tabGempa 
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}
	for i := 0; i < n; i++ {
        fmt.Printf("%s %d %.1f %s\n",
            T[i].nama,
            T[i].tahun,
            T[i].magnitudo,
            kategori(T[i].magnitudo))
    }
	rata := rataMagnitudo(T, n)
    fmt.Printf("Rata-rata magnitudo: %.2f\n", rata)

	idx := idxTerkuat(T, n)

    fmt.Printf("Gempa terkuat: %s %d %.1f\n",
        T[idx].nama,
        T[idx].tahun,
        T[idx].magnitudo)

	besar := hitungKategori(T, n, "Besar")
    fmt.Printf("Jumlah gempa kategori Besar: %d\n", besar)

}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/10-fungsi/unguided_1/gempa_output.png

#### Deskripsi
Program dibuat untuk mengolah data gempa bumi berdasarkan nama wilayah, tahun kejadian, dan nilai magnitudo. Program menentukan kategori setiap gempa berdasarkan magnitudonya, menghitung rata-rata magnitudo, mencari gempa dengan magnitudo tertinggi, serta menghitung jumlah gempa yang termasuk kategori Besar. Setiap proses perhitungan dilakukan menggunakan fungsi sesuai ketentuan soal.

## Kesimpulan
fungsi dalam bahasa Go dapat membantu membuat program lebih terstruktur dan mudah dipahami. Fungsi digunakan untuk mengolah data, melakukan perhitungan, menentukan nilai tertentu, serta mengembalikan hasil menggunakan return.

## Referensi
1. Learn Go. (n.d.). Functions. https://www.learn-golang.org/en/Functions
2. The Go Authors. (n.d.). A Tour of Go. https://go.dev/tour/