### Tugas Minggu 9
#### Maulana Nur Mafaza
#### 109092600015

##### 1. Soal Konversi Jarak
#### Kode

```go
package main

import "fmt"

func konversi(km, m, cm int) int {
    return km*100000 + m*100 + cm
}

func main() {
    var km, m, cm int
    fmt.Scan(&km, &m, &cm)

    hasil := konversi(km, m, cm)
    fmt.Print("Hasil konversi = ", hasil, "cm")
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/09-fungsi/konversiJarak_output.png

##### 2.Soal Kuadran
#### Kode

```go
package main

import "fmt"

func kuadran (x, y int) string {
	switch{
	case x == 0 && y == 0:
		return "Titik Asal"
	case x == 0:
		return "Titik pada sumbu Y"
	case y == 0:
		return "Titik pada sumbu X"
	case x > 0 && y > 0:
		return "Kuadran I"
	case x < 0 && y > 0:
		return "Kuadran II"
	case x < 0 && y < 0:
		return "Kuadaran III"
	case x > 0 && y < 0:
		return "Kuadran IV"
	default:
		return "Tidak valid"

	}
}
func main (){
	var x, y int
	fmt.Scan(&x, &y)

	hasil := kuadran (x, y)
	fmt.Println(hasil)
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/09-fungsi/kuadran_output.png

##### 3. Soal Tarif Parkir
#### Kode

```go
package main

import "fmt"

func tarif(jam int) int{
	if jam <= 2{
		return 5000
	} else {
	return 5000 + (jam - 2) * 3000
	}
}
func main(){
	 var jam int
	 
	 fmt.Println("Lama parkir: ")
	 fmt.Scan(&jam)

	 biaya := tarif(jam)
	 fmt.Println("Total tarif parkir = Rp", biaya)
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/09-fungsi/tarifParkir_output.png

##### 4. Soal Deret Kuadrat
#### Kode

```go
ackage main

import "fmt"

func deretKuadrat(n int) int {
    total := 0

    for i := 1; i <= n; i++ {
        total = total + i*i
    }
    return total
}

func main() {
    var n int
    fmt.Scan(&n)

    hasil := deretKuadrat(n)
    fmt.Println("Hasil penjumlahan =", hasil)
}

```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/09-fungsi/deretKuadrat_output.png

##### 5.Soal Kombinasi
#### Kode

```go
package main

import "fmt"

func faktorial(n int) int {
    hasil := 1

    for i := 1; i <= n; i++ {
        hasil = hasil * i
    }
    return hasil
}

func kombinasi(x, y int) int {
    if x < y {
        x, y = y, x
    }
    return faktorial(x) / (faktorial(y) * faktorial(x-y))
}
func main() {
    var x, y int
    fmt.Scan(&x, &y)

    fmt.Println("Faktorial X =", faktorial(x))
    fmt.Println("Faktorial Y =", faktorial(y))
    fmt.Println("Kombinasi =", kombinasi(x, y))
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/09-fungsi/kombinasi_output.png