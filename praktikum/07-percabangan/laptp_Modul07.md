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

#### Deskripsi

### 3. hari.go

```go

```

##### Output


#### Deskripsi


### 4. terbesar.go

```go

```

##### Output


#### Deskripsi


## Kesimpulan
