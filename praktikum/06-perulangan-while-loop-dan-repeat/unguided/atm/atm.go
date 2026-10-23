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