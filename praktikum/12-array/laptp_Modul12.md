# <h1 align="center">Tugas Pendahuluan Modul 12 - array </h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

### 1. pelamar

```go
package main

import "fmt"

const NMAX = 100

type applicant struct {
	id string
	name string
	testGrade float64
	testDuration int
}

func salinLolos (A [NMAX]applicant, n int, batas float64) ([NMAX]applicant, int) {
	var B [NMAX]applicant
	m := 0

	for i := 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[m] = A[i]
			m++
		}
	}
	return B, m
}

func rataDurasi (B [NMAX]applicant, m int) float64 {
	if m == 0 {
		return 0
	}
	jumlah := 0

	for i := 0; i < m; i++ {
		jumlah +=B[i].testDuration
	}
	return float64(jumlah) / float64(m)
}

func main () {
	var A[NMAX]applicant
	var n int
	var batas float64

	for n < NMAX {
        fmt.Scan(&A[n].id)

        if A[n].id == "END" {
            break
        }

        fmt.Scan(&A[n].name, &A[n].testGrade, &A[n].testDuration)
        n++
    }

    fmt.Scan(&batas)

    fmt.Println("Data awal:")
    for i := 0; i < n; i++ {
        fmt.Printf("%s %s %.1f %d\n",
            A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
    }

    B, m := salinLolos(A, n, batas)

    fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)

    for i := 0; i < m; i++ {
        fmt.Printf("%s %s %.1f %d\n",
            B[i].id, B[i].name, B[i].testGrade, B[i].testDuration)
    }

    if m == 0 {
        fmt.Println("Tidak ada pelamar yang lolos")
    } else {
        fmt.Printf("Jumlah pelamar lolos: %d\n", m)
        fmt.Printf("Rata-rata durasi: %.2f menit\n",
            rataDurasi(B, m))
    }
}
```

##### Output



#### Deskripsi
Soal ini membahas pengolahan data pelamar kerja menggunakan array dan tipe bentukan. Program menerima data pelamar berupa ID, nama, nilai tes, dan durasi pengerjaan. Data pelamar yang memenuhi batas kelulusan disalin ke array kedua. Selanjutnya, program menampilkan data awal, daftar pelamar yang lolos, jumlah pelamar yang lolos, serta rata-rata durasi pengerjaannya.

### 2. nama

```go
package main

import "fmt"

const MAX = 9999

type arrString [MAX]string

func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(A[i])
	}
	fmt.Println()
}

func sisipData(A *arrString, n *int, p int, x string) {
	for i := *n; i >= p; i-- {
		A[i] = A[i-1]
	}

	A[p-1] = x
	*n++
}

func hapusData(A *arrString, n *int, q int) {
	for i := q - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}

	*n--
}

func main() {
	var A arrString
	var n, p, q int
	var x string

	fmt.Scan(&n)
	bacaData(&A, n)

	fmt.Scan(&p, &x)
	fmt.Scan(&q)

	fmt.Print("Data awal: ")
	cetakData(A, n)

	sisipData(&A, &n, p, x)

	fmt.Print("Setelah penyisipan: ")
	cetakData(A, n)

	hapusData(&A, &n, q)

	fmt.Print("Setelah penghapusan: ")
	cetakData(A, n)
}
```

##### Output



#### Deskripsi
Soal ini membahas operasi penyisipan dan penghapusan data pada array. Program menerima sejumlah nama, kemudian menyisipkan nama baru pada posisi yang ditentukan dan menghapus nama pada posisi tertentu setelah penyisipan. Setiap operasi dilakukan dengan menggeser elemen array agar urutan data tetap terjaga. Program menampilkan isi array sebelum penyisipan, setelah penyisipan, dan setelah penghapusan.

## Kesimpulan
array dalam bahasa pemrograman Go dapat digunakan untuk menyimpan dan mengolah kumpulan data secara terstruktur. Pada soal pertama, array digunakan untuk menyeleksi data pelamar berdasarkan nilai tes serta menghitung jumlah dan rata-rata durasi pengerjaan pelamar yang lolos. Pada soal kedua, array digunakan untuk melakukan penyisipan dan penghapusan nama dengan menggeser elemen agar urutan data tetap terjaga.