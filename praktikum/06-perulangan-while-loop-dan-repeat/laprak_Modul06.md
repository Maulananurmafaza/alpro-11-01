# <h1 align="center">Laporan Praktikum Modul 06 - PERULANGAN WHILE</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

## Dasar Teori

### A. Perulangan While di Go
Dalam bahasa Go tidak ada while secara langsung seprti pada beberapa bahasa pemrograman lainya, sebagai gantinya fungsi while bisa kita ganti dengan menggunakan for yang diikuti suatu kondisi.Perulangan akan terus berjalan selama kondisi bernilai true dan berhenti jika kondisi bernilai false
### B. Kondisi Berhenti dan Penelusuran String
Kondisi berhenti adalah kondisi dimana untuk mengahiri perulangan ketika syarat tertentu terpenuhi. Jika string adalah proses memeriksa karakter dalam teks untuk menemukan/menghitung karakter tertentu.

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
 Soal ini menggunakan perulangan for dalam Go untuk mencari bilangan terkecil yang pangkat tiganya lebih besar dari nilai batas yang dimasukkan. Nilai n bertambah hingga kondisi terpenuhi, kemudian hasilnya ditampilkan.


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
Soal ini menggunakan perulangan for dalam Go untuk mencari bilangan terbesar yang kuadratnya masih lebih kecil dari nilai batas yang dimasukkan. Nilai n bertambah hingga kondisi tidak terpenuhi, kemudian hasilnya ditampilkan.
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
Program ini menggunakan bahasa Go untuk menghitung jumlah huruf vokal (a, i, u, e, o) dalam sebuah kata atau kalimat. Program menggunakan perulangan for untuk memeriksa setiap karakter dan menampilkan jumlah huruf vokal yang ditemukan.
## Unguided

### 1. atm

```go
package main

import "fmt"

func main() {
    var password string
    var pilihan int
    var saldo int = 10000000
    var penarikan int
    kesempatan := 0

    for kesempatan < 3 {
        fmt.Print("Masukkan password ATM: ")
        fmt.Scan(&password)

        if password == "1234" {
            fmt.Println("Login berhasil!")
            break
        } else {
            kesempatan++
            fmt.Println("Password salah!")

            if kesempatan == 3 {
                fmt.Println("ATM terblokir.")
                return
            }
        }
    }

    for {
        fmt.Println("\n=== MENU ATM ===")
        fmt.Println("1. Lihat Saldo")
        fmt.Println("2. Tarik Uang")
        fmt.Println("3. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            fmt.Println("Saldo Anda: Rp", saldo)
        case 2:
            fmt.Print("Masukkan jumlah penarikan: Rp ")
            fmt.Scan(&penarikan)

            if penarikan > 0 && penarikan <= saldo {
                saldo -= penarikan
                fmt.Println("Penarikan berhasil!")
                fmt.Println("Sisa saldo: Rp", saldo)
            } else {
                fmt.Println("Saldo tidak cukup atau jumlah tidak valid.")
            }
        case 3:
            fmt.Println("Terima kasih telah menggunakan ATM.")
            return

        default:
            fmt.Println("Pilihan tidak tersedia.")
        }
    }
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/praktikum/06-perulangan-while-loop-dan-repeat/unguided/atm/unguided_atm_output.png

#### Deskripsi
Program dirancang untuk meminta pengguna memasukkan password dengan maksimal tiga kali kesempatan. Jika password benar, pengguna dapat mengakses menu ATM untuk melihat saldo, menarik uang, atau keluar dari program. Jika password salah sebanyak tiga kali, ATM akan terblokir dan program berhenti. Saldo awal yang digunakan adalah Rp10.000.000.

## Kesimpulan
perulangan while dalam bahasa Go dapat diterapkan menggunakan for dengan kondisi tertentu. Perulangan ini digunakan untuk menyelesaikan berbagai masalah, seperti mencari bilangan terkecil, bilangan terbesar, menghitung huruf vokal, dan membuat simulasi ATM. Dengan memahami perulangan dan kondisi berhenti, program dapat menjalankan perintah secara berulang hingga kondisi yang ditentukan terpenuhi.
## Referensi
1. Go Programming Language. The Go Programming Language Specification — For Statements. Dokumentasi resmi Go
2. Go.dev. A Tour of Go — Flow Control Statements: For. A Tour of Go
3. Go.dev. A Tour of Go — Strings, bytes, and runes. Dokumentasi string Go