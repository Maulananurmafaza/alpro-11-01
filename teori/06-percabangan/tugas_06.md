### Tugas Minggu 6
#### Maulana Nur Mafaza
#### 109092600015

##### 1. Profit
#### Kode

```go
package main

import"fmt"

func main(){
	var bulanSebelumnya, bulanIni float64
	fmt.Scan (&bulanSebelumnya, &bulanIni)

	if bulanIni > bulanSebelumnya {
		fmt.Println("Peningkatan sebesar", bulanIni - bulanSebelumnya)
	} else if bulanIni < bulanSebelumnya {
		fmt.Println("Penurunan sebesar", bulanSebelumnya - bulanIni)
	} else {
		fmt.Println("Tetap")
	}
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/06-percabangan/profit_output.png

##### 2. sepakBola
#### Kode

```go
package main

import "fmt"

func main() {
	var gol1, gol2, gol3, gol4 int

	fmt.Print("Masukkan jumlah gol 4 tim (pisahkan dengan spasi): ")
	_, err := fmt.Scan(&gol1, &gol2, &gol3, &gol4)
	if err != nil {
		fmt.Println("Masukan tidak valid.")
		return
	}
	terbanyak := gol1
	if gol2 > terbanyak {
		terbanyak = gol2
	}
	if gol3 > terbanyak {
		terbanyak = gol3
	}
	if gol4 > terbanyak {
		terbanyak = gol4
	}

	tersedikit := gol1
	if gol2 < tersedikit {
		tersedikit = gol2
	}
	if gol3 < tersedikit {
		tersedikit = gol3
	}
	if gol4 < tersedikit {
		tersedikit = gol4
	}

	fmt.Printf("Jumlah gol terbanyak: %d\n", terbanyak)
	fmt.Printf("Jumlah gol tersedikit: %d\n", tersedikit)
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/06-percabangan/sepakBola_output.png

##### 3. Parkir
#### Kode

```go
package main

import "fmt"

func main(){
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1, &h2, &m2)

	mulai := h1*60 + m1
	selesai := h2*60 + m2

	if h2 < h1 {
		selesai += 12*60
	}

	durasi := selesai - mulai

	hh := durasi / 60
	mm := durasi % 60

	fmt.Println(hh, "jam", mm, "menit")
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/06-percabangan/parkir_output.png

##### 4. akhirTahun
#### Kode

```go
package main

import "fmt"

func main(){
	var belanja int
	var bersedia bool
	fmt.Scan(&belanja, &bersedia)

	kartu := bersedia
	diskon := belanja >= 100000
	cashback := belanja >= 200000 && kartu

	bayar := belanja

	if diskon {
		bayar = bayar - (belanja * 10 / 100)
	}
	if cashback {
		bayar = bayar - 75000
	}
	fmt.Println("kartu?", kartu)
	fmt.Println("diskon?", diskon)
	fmt.Println("cashback?", cashback)
	fmt.Println("Rp", bayar)

}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/06-percabangan/akhirTahun_output.png

##### 5. Avatar
#### Kode

```go
package main

import "fmt"

func main() {
    var p int
    fmt.Scan(&p)

    appaDewasa := p / 5
    if appaDewasa > 3 {
        appaDewasa = 3
    }

    sisa := p - appaDewasa*5

    if sisa > 0 && appaDewasa < 3 {
        appaDewasa++
        sisa = p - appaDewasa*5
        if sisa < 0 {
            sisa = 0
        }
    }

    appaKecil := sisa / 2
    if appaKecil > 5 {
        appaKecil = 5
    }

    sisa = sisa - appaKecil*2

    fmt.Print("Dewasa ", appaDewasa)

    if appaKecil > 0 {
        fmt.Print(", kecil ", appaKecil)
    }

    if sisa > 0 {
        fmt.Print(", dan ", sisa, " tak berangkat")
    }

    fmt.Println()
}
```
##### output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/06-percabangan/avatar_output.png