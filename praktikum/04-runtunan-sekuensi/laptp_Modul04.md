# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtutan dan Sekuensi</h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

```go
package main

import "fmt"

func main(){
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if intNum > 0 {
		fmt.Println(intNum > 0)
	}
	if intNum >= 5 && intOther < 11 {
		fmt.Println(intNum >= 5 && intOther < 11)
	}
	if sngNum > 0 || (intNum > 0 && -1 * intOther == -10) {
		fmt.Println(sngNum > 0 || (intNum > 0 && -1 * intOther == -10))
	}
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/evaluasiEkspresiKontrol/pendahuluan_evaluasiEkspresiKontrol_output.png


#### Deskripsi
membahas evaluasi ekspresi kontrol dalam bahasa Go, yaitu menentukan apakah suatu kondisi menghasilkan nilai true atau false

### 2.Tracing: Evaluasi Pernyataan Kondidsi

```go
package main

import "fmt"

func main(){
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
}

if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/evaluasiPernyataanKondisi/pendahuluan_evaluasiPernyataanKondisi_output.png

#### Deskripsi
membahas evaluasi pernyataan kondisi dalam bahasa Go dengan menggunakan percabangan if, else if, dan else

### 3. Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Println("Masukkan tahun: ")
	fmt.Scan(&tahun)
	fmt.Println("Masukkan bulan: ")
	fmt.Scan(&bulan)
	kabisat := (tahun % 400 == 0) || (tahun % 4 == 0 && tahun % 100 != 0)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println("Jumlah hari: 31")
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println("Jumlah hari: 30")
	case "Feb":
		if kabisat {
			fmt.Println("Jumlah hari: 29")
		} else {
			fmt.Println("Jumlah hari: 28")
		}
	default:
		fmt.Println("_")
	}
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/jumlahHari/Pendahuluan_jumlahHari_output.png

#### Deskripsi
membahas pembuatan program dalam bahasa Go untuk menentukan jumlah hari dalam suatu bulan berdasarkan tahun dan nama bulan yang diinput pengguna, termasuk memperhatikan tahun kabisat dan validasi nama bulan


### 4. Switch Case

```go
package main

import "fmt"

func main() {
    var pilihan int

    fmt.Println("1. Air Mineral")
	 fmt.Println("2. Teh")
    fmt.Println("3. Kopi")
    fmt.Print("Pilih minuman (1-3): ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        fmt.Println("Anda memilih Air Mineral")
    case 2:
        fmt.Println("Anda memilih Teh")
	 case 3:
        fmt.Println("Anda memilih Kopi")
    default:
        fmt.Println("Pilihan tidak tersedia")
	}
}
##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/switchCase/pendahuluan_swithCase_output.png

#### Deskripsi
Program ini menggunakan switch case untuk menampilkan jenis minuman berdasarkan pilihan pengguna. Struktur ini membuat program sederhana dan mudah dipahami


## Kesimpulan
Berdasarkan seluruh soal di atas, dapat disimpulkan bahwa bahasa pemrograman Go memiliki berbagai fitur untuk mengatur alur program, seperti operator logika, percabangan if-else, dan switch case. Melalui soal-soal tersebut, kita belajar mengevaluasi kondisi true atau false, menelusuri alur eksekusi program, menghitung nilai akhir variabel, serta membuat program untuk menentukan jumlah hari dalam suatu bulan dengan memperhatikan tahun kabisat dan validasi input.