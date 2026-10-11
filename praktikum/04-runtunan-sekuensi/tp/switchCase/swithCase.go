package main

import "fmt"

func main() {
    var pilihan int

    fmt.Println("1. Air Mineral")
	 fmt.Println("2. Teh")
    fmt.Println("3. Kopi")
    fmt.Print("Pilih minuman (1-3): ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        fmt.Println("Anda memilih Air Mineral")
    case 2:
        fmt.Println("Anda memilih Teh")
	 case 3:
        fmt.Println("Anda memilih Kopi")
    default:
        fmt.Println("Pilihan tidak tersedia")
	}
}