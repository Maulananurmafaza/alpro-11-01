# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan While </h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. lulus.go

```go
package main
import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&nilai)

		if nilai < 0 {
			fmt.Println("Program selesai.")
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian")
		} else {
			fmt.Println("Anda tidak lulus ujian")
		}
	}
}
```

##### Output



#### Deskripsi
Prigram mengecek kelulusan ujian secara berulang berdasarkan nilai integer yang dimasukkan pengguna. Program terus meminta input hingga pengguna memasukkan angka negatif (kurang dari 0) yang berfungsi sebagai kondisi berhenti. Setiap kali nilai valid dimasukkan, program mengevaluasinya: jika nilai 60 atau lebih, program menampilkan pesan "Anda lulus ujian", dan jika kurang dari 60, program menampilkan "Anda tidak lulus ujian".

### 2. deret.go

```go
package main

import "fmt"

func main(){
	var n, i, s int
	fmt.Println("Masukkan bilangan positif n: ")
	fmt.Scan(&n)
	i = 1
	s = 0

	for i <= n {
		s = s + i
		i++
	}
	fmt.Println("jumlah deret=", s)
}
```

##### Output



#### Deskripsi
Program menghitung jumlah total dari deret perulangan bilangan bulat positif dari 1 sampai n

### 3. duaTerendah.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	siswa1, siswa2 := "", ""
	nilai1, nilai2 :=  999, 999

	for i := 0; i < n; i++ {
		var siswa string
		var nilai int
		fmt.Scan(&siswa, &nilai)

		if nilai < nilai1 {
			nilai2 = nilai1
			siswa2 = siswa1
			nilai1 = nilai
			siswa1 =  siswa
		} else if nilai < nilai2 {
			nilai2 = nilai
			siswa2 = siswa
		}
	}
	fmt.Println(siswa1)
	fmt.Println(siswa2)
}
```

##### Output


#### Deskripsi
Program membaca n, lalu melakukan perulangan untuk menerima masukan nama dan nilai siswa satu per satu. Program melacak dua nilai terkecil (nilai1 dan nilai2) beserta nama siswanya (siswa1 dan siswa2). Setelah semua data selesai diproses, program mencetak nama kedua siswa dengan nilai terendah tersebut secara berurutan.

### 4. biaya.go

```go
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0

	for i := 2; i <=n; i++ {
		perTahun := biaya * 0.06
		biaya += perTahun
	
	}
	fmt.Printf("%.0f\n", biaya)
}
```

##### Output


#### Deskripsi
Program menerima masukan jumlah tahun (n), menetapkan biaya awal sebesar 10.000.000, lalu melakukan perulangan untuk menambahkan kenaikan biaya tersebut setiap tahunnya hingga tahun ke-n.

## Kesimpulan
keempat program Go di atas merupakan contoh penerapan logika pemrograman dasar perulangan while dengan memggunakan  for untuk menyelesaikan kasus-kasus komputasi sederhana
