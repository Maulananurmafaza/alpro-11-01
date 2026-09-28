# <h1 align="center">Laporan Praktikum Modul [02] - [Bahasa Pemrograman Go]</h1>
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

### 1. [skor.go]

```go
package main

import "fmt"

func main(){
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	// Menghitung total & rata-rata (pembagian bilangan)
	total:= skorMatematika + skorBahasaInggris
	ratarata := total/2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(ratarata)
}
```
#### Deskripsi
Membuat program yang membaca input nama seorang siswa, skor matematika, dan skor bahasa inggris. Kemudian program menghitung total gabungan dari kedua nilai, lalu mencari rata-ratanya dari total nilai tadi. Bagian guided yang mengimplementasikan yaitu logika dibalik program syntax pemrograman. Hail yang diperoleh adalah output nama siswa, total nilai matematika dan bahasa inggris serta rata-ratanya. 


### 2. [tukar.go]

```go
package main

import "fmt"

func main() {
	var a, b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a, b = b, a

	//output
	fmt.Println(a)
	fmt.Println(b)
}

```
#### Deskripsi
Membuat program yang mampu membaca dua input bilangan bulat lalu menukar nilai kedua bilangan bulat tersebut. Bagia guided yang mengimplementasikan yaitu logika serta syntax pemrograman. Hasil yang diperoleh berupa output yang berkebalikan dari input yang diberikan, misalnya input adalah a dan b, maka outputnya adalah b dan a.



## Unguided

### 1. [cacahuang.go]

```go
package main

import "fmt"

func main(){
	var uang int
	fmt.Println("Msukan nominal uang")
	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seRibu := sisa / 1000

	fmt.Printf("Sepuluh Ribu = %d, Lima Ribu = %d, Seribu = %d\n", sepuluhRibu, limaRibu, seRibu)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/unguided_cacahuang_output.png)


#### Deskripsi
Membuat program yang berfungsi menguraikan nominal uang menjadi pecahan (hanya bilangan bulat) 10.000, 5.000, dan 1.000. Bagian guided yang diimplementasikan yakni cara menggunakan variabel integer, fmt.PrintIn, fmt. Scan, dan algoritma pembagian dan modulo. Bagian unguided yang mengimpletasikan yakni cara menggunakan format specifier seperti %d, dan fmt.Printf. Hasil yang diperoleh berupa output berapa banyak uang nominal 10.000, 5.000, dan 1.000 dari input total 
### 2. [kalkulator.go]

```go
package main

import "fmt"

func main() {
	var a,b int

	//membaca input
	fmt.Print("masukan nilai a")
	fmt.Scan(&a)
	fmt.Print("masukan nilai b")
	fmt.Scan(&b)

	//menampilkan output
	fmt.Print("Hasil Penjumlahan: ")
	fmt.Println(a + b)
	fmt.Print("Hasil Pengurangan: ")
	fmt.Println(a - b)
	fmt.Print("Hasil Perkalian: ")
	fmt.Println(a * b)
	fmt.Print("Hasil Pembagian: ")
	fmt.Println(a / b)
	fmt.Print("Hasil Sisa Hasil Bagi: ")
	fmt.Println(a % b)

}
```

##### Output
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/kalkulator/unguided_kalkulator_output.png)

#### Deskripsi
Membuat program yang menghitung hasil penjumlahan, pengurangan, perkalian, pembagian, dan modulo dari dua input yang dimasukkan. Bagian guided yang diimplementasikan adalah cara menggunakan variabel integer, fmt.scan, fmt.Println, serta logika matematika dasar. Bagian unguided yang mengimpelemntasikan adalah cara penggunaan format specifier seperti %d yang mana digunakan untuk menandai tempat yang akan diisi oleh nilai variabel nantinya, dan fmt.Printf untuk mencetak teks dengan format, yang dalam program ini berupa format specifier %d. Hasilnya berupa output hasil operasi matematika sederhana dari dua input yang dimasukkan.

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
