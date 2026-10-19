# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For</h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

## Dasar Teori

### A. Perulangan For di Go
perulangan (looping) adalah mekanisme untuk mengeksekusi sekumpulan kode secara berulang selama kondisi tertentu terpenuhi.Go hanya memiliki satu konstruksi perulangan, yaitu kata kunci for. Go tidak memiliki while atau do-while.

### B. For hanya kondisi dan Perulangan Bersarang
Karena Go tidak memiliki kata kunci while, Anda bisa menggunakan for dengan hanya menuliskan kondisinya saja.
Sedangkan Perulangan bersarang adalah perulangan di dalam perulangan. Biasanya digunakan untuk mengolah data multidimensi.

## Guided

### tabungan.go

```go
package main

import "fmt"

func main(){
	var a, b, n int
	
	fmt.Scan(&a, &b, &n)

	total :=0;
	tabunganHariIni := a;

	for i := 1; i <= n; i++{
		total = total + tabunganHariIni

		tabunganHariIni = tabunganHariIni + b
	}

	fmt.Println(total)
}
```
#### Deskripsi
 

### 2. terbesar.go

```go
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}
```
#### Deskripsi


### 3. bioskop.go

```go
package main

import "fmt"

func main(){
	var t int
	fmt.Print("input t (jumlah transaksi yang akan dilakukan): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++{
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if jenisFilm == "Reguler" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 40000
			} else{
				hargaSatuan = 50000
			}
		} else if jenisFilm == "3D" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 60000
			} else{
				hargaSatuan = 70000
			}
		} else if jenisFilm == "IMAX" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 80000
			} else{
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}

	fmt.Println("output total pendapatan:", totalPendapatan)

}
```
#### Deskripsi


## Unguided

### 1. prima

```go
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <=n; i++{
		if n%i == 0 {
			if jumlah > 0 {
				fmt.Print(" ")
			}
			fmt.Print(i)
			jumlah++
		}
	}
	fmt.Println()
	if jumlah == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}

}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/prima/unguided_prima_output.png

#### Deskripsi


### 2. ecct

```go
package main

import "fmt"

func main(){
	var m int 
	fmt.Scan(&m)
	lulus := 0

	for i := 1; i <= m; i++{
		var P, V, GA, FS, CI, TC int
		fmt.Scan(&P, &V, &GA, &FS, &CI, &TC)

		total := P + V + GA + FS + CI + TC
		rataRata := float64(total) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}
	fmt.Println(lulus)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/unguided/ecct/unguided_ecct_output.png

#### Deskripsi


## Kesimpulan


## Referensi
