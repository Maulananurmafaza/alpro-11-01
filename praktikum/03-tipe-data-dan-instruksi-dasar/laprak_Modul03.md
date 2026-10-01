# <h1 align="center">Laporan Praktikum Modul 03 - Variabel dan Operator</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Variabel dan Tipe Data di Go
Variabel merupakan tempat penyimpanan dalam memori komputer untuk menampung data selama program berjalan.Yang dimana, bahasa Go menggunakan sistem statically typed, yang artinya setiap variabel wajib memiliki tipe data yang pasti sejak awal dan tidak bisa diisi sembarang tipe lain ditengah proses
### B. Operator Aritmatika dan Assignments

#### 1. Operator Aritmatika dan Modulo
Operator aritmatika berfungsi untuk melakukan perhitungan seperti +, -, *, /, dan pencarian sisa hasil bagi.Operator ini biasanya digunakan dalam beragai masalah seperti menghitung total belanja, menentukan rata-rata nilai,dll.Sedangkan modulo biasanya digunakan untuk untuk memperoleh sisa hasil pembagian dua bilangan bulat
#### 2. Pertukaran Nilai (Multiple Assignment)
Pertukaran nilai digunakan setika suatu program perlu memindahkan nilai dari suatu variabel ke variabel lainnya tanpa mengubah nilai yang dipertukarkan secara tidak sengaja.
## Guided

### 1. [kasir.go]

```go
package main

import "fmt"

func main(){
	var x int 
	fmt.Println("Masukan nominal")
	fmt.Scan(&x)

	var sepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seRibuan int = sisa / 1000

	fmt.Println(sepuluhRibuan, limaRibuan, seRibuan)
}
```
#### Deskripsi
 membahas pembuatan program dalam bahasa pemrograman Go untuk menghitung jumlah lembar uang yang harus diberikan sebagai kembalian kepada pembeli,mengubah nilai uang kembalian menjadi jumlah lembar uang berdasarkan tiga pecahan yang tersedia, yaitu Rp10.000, Rp5.000, dan Rp1.000. Soal ini melatih penggunaan variabel, operasi pembagian, modulus, dan output pada bahasa Go.


### 2. [konversi.go]

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}

```
#### Deskripsi
meminta kita untuk membuat sebuah program menggunakan bahasa pemrograman Go yang berfungsi untuk mengonversi suhu dari derajat Celsius menjadi Kelvin ,merupakan latihan dasar pemrograman Go yang berfokus pada input, variabel bertipe bilangan real, operasi aritmatika penjumlahan, dan output. Program cukup membaca nilai Celsius, menambahkan 273 sesuai rumus yang diberikan, lalu menampilkan hasilnya dalam Kelvin



## Unguided

### 1. [konversi_hari.go]

```go
package main

import "fmt"

func main(){
	var hari int
	var tahun, bulan, minggu, sisaHari int

	fmt.Println("masukan jumlah Hari")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari = hari % 360
	bulan = hari / 30 
	hari = hari % 30
	minggu = hari / 7
	sisaHari = hari % 7

	fmt.Println("tahun", tahun)
	fmt.Println("bulan", bulan)
	fmt.Println("minggu", minggu)
	fmt.Println("sisaHari", sisaHari)


}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_hari/unguided_konversi_hari_output.png)


#### Deskripsi
Meminta kita mengkonversi jumlah hari menjadi satuan tahun, bulan, minggu dan sisa hari.Program kemudian harus mengambil total hari sebagai input,kemudian secara bertahap mengubahnya menjadi tahun-bulan-minggu-sisahari
### 2. [konversi_suhu.go]

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu celcius")
	fmt.Scan(&celcius)

	reamur := celcius * 4/5
	fmt.Println("konversi ke Suhu reamur", reamur)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_suhu/unguided_konversi_suhu_output.png)

#### Deskripsi
Meminta kita untuk mengonversi suhu dari derajat celcius ke reamur,program menerima satu nilai suhu celsius sebagai input, kemudian mengubahnya menjadi suhu reamur menggunakan rumus R = (4/5) × C,dari soal ini kita dilatih untuk memahami hubungan antara input → proses/perhitungan → output serta penggunaan tipe data numerik

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Dari praktikum kali ini kita dapat memahami penerapan variabel, tipe data numerik, operator aritmatika, input, proses, dan output dalam bahasa Go. Pemahaman ini menjadi dasar penting untuk mengembangkan kemampuan berpikir logis, menyusun algoritma secara sistematis, dan membuat program yang lebih kompleks.

## Referensi
1. The Go Authors. The Go Programming Language Specification. Dpak diakses melalui
tautan https://go.dev/ref/spec
2. The Go Authors. (n.d.). A Tour of Go: Variables with Initializers. https://go.dev/tour/basics/9