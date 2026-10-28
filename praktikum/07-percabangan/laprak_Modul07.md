# <h1 align="center">Laporan Praktikum Modul 07 - Algoritma dan Pemrograman Percabangan</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Algoritma

### B. Pemrograman Percabangan


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


## Unguided

### 1. besok.go

```go

```

##### Output

#### Deskripsi


### 2. besok.go

```go

```

##### Output

#### Deskripsi

## Kesimpulan

## Referensi
