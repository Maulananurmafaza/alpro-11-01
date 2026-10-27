# <h1 align="center">Tugas Pendahuluan Modul 07 - Percabangan </h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. Perbedaan if, if-else, else if, Switch pada go dan java

- if
if adalah percabangan dalam bahasa Go yang digunakan untuk menjalankan kode jika suatu kondisi bernilai benar (true).

- if-else
if else adalah percabangan yang digunakan untuk memilih satu dari dua tindakan berdasarkan kondisi. Jika kondisi if benar, perintah if dijalankan. Jika salah, perintah else dijalankan.

- else if
else if adalah percabangan yang digunakan untuk memeriksa beberapa kondisi secara berurutan sampai ditemukan kondisi yang benar.

- Switch
Switch adalah percabangan dalam bahasa Go yang digunakan untuk memilih satu perintah dari beberapa pilihan berdasarkan nilai suatu variabel atau ekspresi.

a. Switch pada go
Pada penulisan contohnya; switch variabel.
Pada tanda kurung contohnya; tidak memerlukan tanda kurung.

b. Switch pada java
Pada penulisan contohnya; switch (variabel).
Pada tanda kurung contohnya; menggunakan tanda kurung pada ekspressi.


### 2. bilangan.go

```go
ackage main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Nol")
	} else {
		p, t := "Ganjil", "Negatif"
		if n%2 == 0 { p = "Genap" }
		if n > 0 { t = "Positif" }
		fmt.Printf("%s, %s\n", p, t)
	}
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/07-percabangan/tp/bilangan/pendahuluan_bilangan_output.png
#### Deskripsi
Soal ini meminta pembuatan program yang mampu menganalisis karakteristik matematis dari sebuah bilangan bulat yang diinputkan.

### 3. hari.go

```go
package main

import "fmt"

func main(){
	var hari int
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu (akhir pekan)")
	case 7:
		fmt.Println("Minggu (akhir pekan)")
	default:
		fmt.Println("Nomor tidak valid")
	
	}
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/07-percabangan/tp/hari/pendahuluan_hari_output.png

#### Deskripsi
Membuat program dalam bahasa Go menggunakan struktur kontrol switch yang menerima masukan nomor hari dari 1 hingga 7.

### 4. terbesar.go

```go
package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	terbesar := a

	if b > terbesar {
		terbesar = b
	}
	if c > terbesar {
		terbesar = c
	}
	fmt.Println(terbesar)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/07-percabangan/tp/terbesar/pendahuluan_terbesar_output.png

#### Deskripsi
Soal ini menuntut pembuatan algoritma perbandingan numerik untuk mencari nilai terbesar dari sekumpulan tiga bilangan bulat yang diberikan.

## Kesimpulan
modul ini melatih kemampuan dalam menerapkan logika percabangan di bahasa Go. Struktur if digunakan untuk pengujian kondisi yang dinamis dan kompleks, sedangkan switch digunakan untuk pemetaan pilihan kondisi yang pasti dan terstruktur.