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
