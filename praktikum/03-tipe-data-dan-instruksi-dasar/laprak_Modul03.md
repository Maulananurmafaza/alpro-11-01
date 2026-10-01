# <h1 align="center">Laporan Praktikum Modul [03] - [Variabel dan Operator]</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
Go atau Golang adalah bahasa pemrograman yang dikembangkan oleh Google. Go dirancang untuk membuat program yang sederhana, cepat, mudah dipelihara, dan mendukung pemrograman konkuren.
### B. [Package dan Struktur Program di Go]

#### 1. [Pengertian Package main dan func main()]
package main digunakan untuk menunjukkan bahwa file Go tersebut termasuk dalam package utama yang dapat menghasilkan program yang bisa dijalankan (executable)
func main() adalah fungsi utama dalam program Go. Ketika program dijalankan, Go akan memulai eksekusi program dari fungsi main().
#### 2. [Tipe Data dan Deklarasi Variabel di Go]
Tipe data dalam Go ada beberapa macam, diantaranya yaitu bilangan bulat, bilangan real, teks, dan boolean. \
a. Bilangan bulat atau integer adalah tipe data yang digunakan untuk menyimpan bilangan bulat, baik positif maupun negatif. Tipe data ini dideklarasikan dengan int, int8, int32, dan int64. \
b. Bilangan real atau float adalah tipe data yang digunakan untuk menyimpan bilangan desimal. Tipe data ini dideklarasikan dengan float32 dan float64. \
c. Teks atau string adalah tipe data yang digunakan untuk menyimpan teks atau karakter. Tipe data ini dideklarasikan dengan tanda petik ganda atau "...". \
d. Boolean adalah tipe data yang hanya memiliki dua kemungkinan, true atau false. Tipe data ini dideklarasikan dengan kata kunci bool. 

Deklarasi variabel di Go dilakukan dengan dua cara, yaitu bisa dengan var dan :=. \
a. Deklarasi var digunakan untuk menyebutkan nama variabel dan tipe data dengan jelas. \
b. Deklarasi := digunakan agar compiler secara otomatis menebak tipe data berdasarkan nilai didalamnya. 

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
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_hari/konversi.go)


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
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_suhu/konversi.go)

#### Deskripsi
Meminta kita untuk mengonversi suhu dari derajat celcius ke reamur,program menerima satu nilai suhu celsius sebagai input, kemudian mengubahnya menjadi suhu reamur menggunakan rumus R = (4/5) × C,dari soal ini kita dilatih untuk memahami hubungan antara input → proses/perhitungan → output serta penggunaan tipe data numerik

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
-Go (Golang) adalah bahasa pemrograman yang sederhana, cepat, efisien, dan mudah dipelihara. \
-Go dikembangkan oleh Google dan mendukung pemrograman konkuren. \
-package main digunakan sebagai package utama dalam program Go. \
-func main() merupakan fungsi utama dan menjadi titik awal eksekusi program. \
-Go memiliki berbagai tipe data, seperti: \
a.int untuk bilangan bulat. \
b.float32 dan float64 untuk bilangan desimal. \
c.string untuk teks. \
d.bool untuk nilai true atau false. \
-Deklarasi variabel dapat dilakukan menggunakan var dan :=. \
a.var digunakan untuk mendeklarasikan variabel dengan tipe data yang dapat ditentukan secara jelas. \
b.:= digunakan untuk mendeklarasikan variabel dengan tipe data yang ditentukan secara otomatis berdasarkan nilainya. \
-Dengan struktur yang sederhana dan tipe data yang jelas, Go dapat digunakan untuk membuat program yang efisien dan mudah dikembangkan.

## Referensi
1.Go Team. (2026). The Go Programming Language Specification. Google LLC. Diakses pada 28 September 2026 melalui https://go.dev/ref/spec. \
2.Go Team. (2026). Standard Library Documentation. Google LLC. Diakses pada 28 September 2026 melalui https://pkg.go.dev/fmt#section-dokumentation.
