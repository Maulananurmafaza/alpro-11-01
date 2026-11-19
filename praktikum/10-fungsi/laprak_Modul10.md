# <h1 align="center">Laporan Praktikum Modul 10 - fungsi</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Fungsi
Fungsi merupakan satu kesatuan rangkaian intruksi yang memberikan atau menghasilkan suatu nilai yang dimana biasanya memetakan input ke suatu nilai yang lain

## Guided

### 1. asas.go

```go
package main

import "fmt"

// Fungsi untuk mengecek tahun kabisat
func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

// Fungsi untuk mengembalikan jumlah hari pada bulan tertentu dalam suatu tahun
func jumlahHari(bulan, tahun int) int {
	var hari int
	if bulan == 2 {
		if kabisat(tahun) {
			hari = 29
		} else {
			hari = 28
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		hari = 30
	} else {
		hari = 31
	}
	return hari
}

// Fungsi untuk menghitung hari ke-berapa dalam tahun tersebut menggunakan perulangan
func hariKe(tanggal, bulan, tahun int) int {
	var total int = 0
	var i int
	for i = 1; i < bulan; i++ {
		total = total + jumlahHari(i, tahun)
	}
	total = total + tanggal
	return total
}

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	// Menghitung urutan hari dalam tahun
	urutanHari := hariKe(tanggal, bulan, tahun)

	// Menentukan jumlah total hari dalam satu tahun (kabisat = 366, biasa = 365)
	var totalHariTahun int
	if kabisat(tahun) {
		totalHariTahun = 366
	} else {
		totalHariTahun = 365
	}

	// Menghitung sisa hari hingga akhir tahun
	sisaHari := totalHariTahun - urutanHari

	// Menampilkan keluaran sesuai format
	fmt.Printf("Hari ke-%d\n", urutanHari)
	fmt.Printf("Sisa hari dalam tahun %d: %d\n", tahun, sisaHari)
}
```
#### Deskripsi
 
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



#### Deskripsi


## Kesimpulan


## Referensi
