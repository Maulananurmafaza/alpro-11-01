# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. Fungsi
Perulangan For berfungsi untuk menjalankan intruksi secara berulang selama kondisi yang ditentukan masih terpenuhi dan program ini membantu programmer membuat kode yang lebih efisien.
# Komponen Perulangan
- Inisiasi

- Kondisi

- post statement


### 3. ganjil

```go
package main

import "fmt"

func main(){
	var n int
	jumlah := 0

	fmt.Println("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil: ")


	for i := 1; i <= n; i+= 2 {
		fmt.Println(i, " ")
		jumlah +=i
	}

	fmt.Println()
	fmt.Println("jumlah =", jumlah)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/ganjil/pendahuluan_ganjil_output.png


#### Deskripsi
Program ini menggunakan bahasa Go perulangan for untuk menampilkan bilangan ganjil dari 1 sampai n dan menghitung jumlahnya.

### 4. Perkalian

```go
package main

import "fmt"

func main(){
	var a, b int
	fmt.Println("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Println("Masukkan batas perkalian b: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/perkalian/pendahuluan_perkalian_output.png


#### Deskripsi
Program ini menggunakan bahasa Go perulangan for untuk menampilkan tabel perkalian suatu bilangan dari 1 sampai batas yang ditentukan

## Kesimpulan
perulangan for dapat digunakan untuk menjalankan instruksi secara berulang. Pada soal bilangan ganjil, perulangan digunakan untuk menampilkan bilangan ganjil dan menghitung jumlahnya. Sedangkan pada soal perkalian, perulangan digunakan untuk menampilkan tabel perkalian sesuai batas yang ditentukan.soal-soal ini membantu memahami penggunaan perulangan, variabel, kondisi, dan operasi matematika dalam pemrograman.