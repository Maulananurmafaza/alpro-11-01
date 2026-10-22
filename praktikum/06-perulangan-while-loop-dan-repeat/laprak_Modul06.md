# <h1 align="center">Laporan Praktikum Modul 06 - PERULANGAN WHILE</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Perulangan While di Go

### B. Kondisi Berhenti dan Penelusuran String


## Guided

### 1. terkecil.go

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
```
#### Deskripsi
 


### 2. terbesar.go 

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}

```
#### Deskripsi

### 3. vokal.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```
#### Deskripsi

## Unguided

### 1. atm

```go

```

##### Output


#### Deskripsi


## Kesimpulan
Dari praktikum kali ini kita dapat memahami penerapan variabel, tipe data numerik, operator aritmatika, input, proses, dan output dalam bahasa Go. Pemahaman ini menjadi dasar penting untuk mengembangkan kemampuan berpikir logis, menyusun algoritma secara sistematis, dan membuat program yang lebih kompleks.

## Referensi
1. The Go Authors. The Go Programming Language Specification. Dpak diakses melalui
tautan https://go.dev/ref/spec
2. The Go Authors. (n.d.). A Tour of Go: Variables with Initializers. https://go.dev/tour/basics/9