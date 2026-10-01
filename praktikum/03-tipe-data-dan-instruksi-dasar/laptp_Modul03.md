# <h1 align="center">Tugas Pendahuluan Modul [03] - Variabel dan Operator Bahasa Pemrograman</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main(){
	var x, y int
	fmt.Println("Jumlah Kue")
	fmt.Scan(&y)
	fmt.Println("Jumlah anggota Keluarga")
	fmt.Scan(&x)

	modulo := y % x
	fmt.Println("Sisa = %d", modulo)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/sisa.go)



#### Deskripsi
[membahas permasalahan pembagian kue secara merata dalam sebuah keluarga dan meminta kita membuat program menggunakan bahasa pemrograman Go untuk menentukan jumlah kue yang masih tersisa setelah seluruh kue dibagikan secara rata kepada setiap anggota keluarga dan menerapkan konsep dasar pemrograman, khususnya operator modulus dalam bahasa pemrograman Go.]

### 2. Konversi

```go
package main

import "fmt"

func main(){
	var x float64

	fmt.Println("Masukkan jarak dalam mil")
	fmt.Scan(&x)

	kilometer := x * 1.6
	fmt.Println("Konversi ke kilometer", kilometer)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/konversi.go)


#### Deskripsi
[Program menerima masukan berupa jarak dalam satuan mil yang dapat berupa bilangan bulat maupun bilangan desimal. Oleh karena itu, nilai masukan disimpan menggunakan tipe data float64. Setelah menerima nilai tersebut, program melakukan proses perhitungan dengan mengalikan jarak dalam mil dengan faktor konversi 1,6.]

## Kesimpulan
[Kedua soal melatih kemampuan untuk menerjemahkan permasalahan sehari-hari menjadi algoritma yang dapat diprogram. Soal pertama berfokus pada operasi modulus untuk mencari sisa pembagian, sedangkan soal kedua berfokus pada operasi aritmatika dan bilangan desimal. Dengan mengerjakan kedua soal, dapat dipahami bagaimana bahasa Go digunakan untuk menerima data dari pengguna, mengolah data menggunakan operator yang sesuai, serta menghasilkan keluaran berdasarkan aturan yang telah ditentukan.]