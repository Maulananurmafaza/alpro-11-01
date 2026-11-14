package main

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
