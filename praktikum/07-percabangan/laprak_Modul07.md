# <h1 align="center">Laporan Praktikum Modul 07 - Percabangan</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Pengertian percabangan
Percabangan adalah struktur dalam pemrograman yang digunakan untuk menentukan keputusan berdasarkan kondisi tertentu. Percabangan memungkinkan program menjalankan perintah yang berbeda sesuai dengan kondisi yang terpenuhi. Percabangan terdiri dari if, if else, else if, dan switch case.
### B. Jenis Percabangan
- if,
if adalah percabangan dalam bahasa Go yang digunakan untuk menjalankan kode jika suatu kondisi bernilai benar (true).

- if-else,
if else adalah percabangan yang digunakan untuk memilih satu dari dua tindakan berdasarkan kondisi. Jika kondisi if benar, perintah if dijalankan. Jika salah, perintah else dijalankan.

- else if,
else if adalah percabangan yang digunakan untuk memeriksa beberapa kondisi secara berurutan sampai ditemukan kondisi yang benar.

- Switch,
Switch adalah percabangan dalam bahasa Go yang digunakan untuk memilih satu perintah dari beberapa pilihan berdasarkan nilai suatu variabel atau ekspresi.


## Guided

### 1. kabisat.go

```go
package main

import "fmt"

func main(){
	var tahun int
	fmt.Scan(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0) || (tahun % 400 == 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```
#### Deskripsi
 Soal meminta program untuk menentukan apakah suatu tahun termasuk tahun kabisat atau bukan berdasarkan aturan pembagian tahun dengan 4, 100, dan 400.

### 2. segitiga.go

```go
package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println ("Bukan segitiga")
	} else if a == b && b == c {
		fmt.Println ("Segitiga sama sisi")
	} else if a == b || a ==c || b == c {
		fmt.Println("Segitiga sama kaki")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}
```
#### Deskripsi
Soal meminta program untuk menentukan jenis segitiga berdasarkan panjang tiga sisinya, yaitu segitiga sama sisi, sama kaki atau sembarang.

### 3. kalkulator.go

```go
package main

import "fmt"

func main(){
	var a, b int
	var op string
	fmt.Scan(&a, &op, &b)

	switch op {
	case "+" :
		fmt.Println(a + b)
	case "-" :
		fmt.Println(a - b)
	case "*" :
		fmt.Println(a * b)
	case "/" :
		if b == 0 {
			fmt.Println ("Error: pembagian dengan nol")
		} else {
			fmt.Printf("%2f\n", float64(a)/float64(b))
		}
	case "%" :
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default :
		fmt.Println("Operator tidak valid")
	}
}
```
#### Deskripsi
Soal meminta program untuk melakuakn operasi mtk sederhana, seperti +, -, *, /, dan % menggunakan switch.

### 4. ongkir.go

```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var berat float64
	var zona int 
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30 kg)")
		return 
	}
	var tarifPerkg int
	switch zona {
	case 1:
		tarifPerkg = 10000
	case 2:
		tarifPerkg = 15000
	case 3:
		tarifPerkg = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}
	beratBulat := int (math.Ceil(berat))
	total := beratBulat * tarifPerkg

	if layanan == "Express" {
		total = total+total/2
	}
	fmt.Println(beratBulat)
	fmt.Println(total)
}
```
#### Deskripsi
Soal meminta program menghitung biaya pengiriman paket berdasarkan berat, zona tujuan dan jenis layanan. Berat paket dibulatkan keatas, layanan express dikenakan biaya tambahan 50% dan paket dengan berat lebih dari 30 kg atau zona tidak valid harus menampilkan pesan kesalahan

## Unguided

### 1. besok.go

```go
package main

import "fmt"

func main() {
    var tanggal, bulan, tahun int
    fmt.Scan(&tanggal, &bulan, &tahun)

    kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)

    if bulan < 1 || bulan > 12 {
        fmt.Println("Tanggal tidak valid")
        return
    }

    hari := 31

    if bulan == 2 {
        hari = 28
        if kabisat {
            hari = 29
        }
    } else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
        hari = 30
    }
    if tanggal < 1 || tanggal > hari {
        fmt.Println("Tanggal tidak valid")
        return
    }
    fmt.Println("Valid")
    tanggal++

    if tanggal > hari {
        tanggal = 1
        bulan++

        if bulan > 12 {
            bulan = 1
            tahun++
        }
    }

    fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}
```

##### Output

#### Deskripsi
Soal meminta program memeriksa validasi tanggal, bulan dan tahun, kemudian menampilkan tanggal untuk besoknya. Program juga harus memperhatikan jumlah hari setiap bulan, tahun kabisat, serta pergantian bulan dan tahun.

### 2. parkir.go

```go
ackage main

import "fmt"

func main(){
	var jenis string
	var jam, tarif int
	fmt.Scan(&jenis, &jam)

	if jenis == "mobil" {
		tarif = 5000 + (jam-1)*3000
		if jam > 24 {
			tarif += 50000
		}
	} else if jenis == "motor" {
		tarif = 3000 + (jam-1)*2000
		if jam > 24 {
			tarif += 30000
		}
	}
	fmt.Println(tarif)
}
```

##### Output

#### Deskripsi
Meminta program menghitung biaya parkir berdasarkan jenis kendaraan dan lam parkir.

## Kesimpulan

## Referensi
