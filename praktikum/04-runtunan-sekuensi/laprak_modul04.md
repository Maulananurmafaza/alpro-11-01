# <h1 align="center">Laporan Praktikum Modul 04 - Runtutan Sekuensi</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Operator Perbandigan dan Logika di GO

#### 1. Operator Perbandingan
Operator seperti >, <, >=, <=, ==, dan != membandingkan dua nilai dan menghasilkan boolean (true atau false).

#### 2. Operator Logika
Operator && (AND) bernilai true jika kedua sisi benar, || (OR) jika salah satu sisi benar, dan ! (NOT) membalik nilai logika 

### B. Percabangan

#### 1. If - Else If - Else
Percabangan ini digunakan untuk mengatur alur program berdasarkan kondisi tertentu agar program dapat mengambil keputusan secara tepat.

#### 2. Switch Case
struktur percabangan yang digunakan untuk memilih dan menjalankan kode berdasarkan nilai yang sesuai dengan pilihan tertentu. switch memeriksa nilai, sedangkan case menentukan pilihan yang akan dijalankan.

## Guided

### 1. grade.go

```go
package main

import (
	"bufio"
	"os"
	
	"fmt"
)
func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scanln(&nilai)
	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```
#### Deskripsi
 Program ini menentukan nilai huruf (A-F) dari nilai angka siswa. Nama dibaca satu baris penuh dengan bufio.Scanner, lalu nilai diperiksa dengan if - else if per rentang: 90-100 A, 80-89 B, 70-79 C, 60-69 D, selebihnya F.

### 2. penilaian.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()

		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		fmt.Println("Keluar dari program.")

	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

```
#### Deskripsi
Program ini menampilkan menu penilaian dan keluar. Pilihan 1 menjalankan penilaian seperti grade.go, pilihan 0 mencetak "Keluar dari program.", dan pilihan lain mencetak "Pilihan tidak valid." lalu berhenti.

### 3. Klasifikasi.go

```go
package main

import (
	"fmt" 
)

func main() {
	var usia int
	var gaji int
	var keterangan string

	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}
	fmt.Println(keterangan)
}
```
#### Deskripsi
Program ini mengelompokkan seseorang berdasarkan usia dan gaji tahunan memakai switch tanpa ekspresi. Case diperiksa berurutan dari atas, dan case pertama yang bernilai true menentukan keterangannya.

## Unguided

### 1. grade(swich)

```go
package main

import (
	"bufio"
	"os"
	"fmt"
)
func main() {
	var nama string
	var nilai float64
	var grade string

	fmt.Print("Nama: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()
	
	fmt.Print("Nilai: ")
	fmt.Scanln(&nilai)

	switch{
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/grade/unguided_grade_output.png

#### Deskripsi
Program ini adalah grade.go yang diubah memakai switch. Switch tanpa ekspresi memeriksa rentang nilai dari tertinggi, dengan F sebagai default.

### 2. pajak

```go
package main

import "fmt"

func main() {
	var penghasilan float64
	var pajak float64

	fmt.Print("Penghasilan dalam juta: ")
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 && penghasilan >= 0 {
		pajak = 0.05 * float64 (penghasilan)
	} else if penghasilan > 50 && penghasilan < 100 {
		pajak = (0.05 * 50) + (0.10 * float64 (penghasilan-50))
	} else if penghasilan >= 100 && penghasilan < 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * float64 (penghasilan-100))
	} else if penghasilan >= 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * float64 (penghasilan-200))
	}
	fmt.Printf("Total pajak yang harus dibayarkan: %.2f juta", pajak)
}
```
##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/pajak/unguided_pajak_output.png

#### Deskripsi
Program ini menghitung pajak penghasilan (juta) dengan if - else if - else. Pajak dihitung bertahap per lapisan: 5%, 10%, 15%, dan 20% sesuai bracket penghasilan.

## Kesimpulan
Praktikum ini melatih pembuatan program yang memilih alur berdasarkan kondisi. If - else if cocok untuk kondisi bertingkat seperti pajak, sedangkan switch membuat pengelompokan seperti nilai huruf lebih rapi.
## Referensi
1. The Go Authors. (n.d.). The Go Programming Language Specification. https://go.dev/ref/spec
2. The Go Authors. (n.d.). A Tour of Go. https://go.dev/tour/
3. The Go Authors. (n.d.). Go Documentation. https://go.dev/doc/

