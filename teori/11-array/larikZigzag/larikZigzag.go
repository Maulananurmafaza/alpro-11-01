package main

import "fmt"

const NMAX = 128

type tabInt [NMAX]int

func isiData(T *tabInt, n int) {
    for i := 0; i < n; i++ {
        fmt.Scan(&T[i])
    }
}

func balik(T *tabInt, n int) {
    for i := 0; i < n/2; i++ {
        T[i], T[n-1-i] = T[n-1-i], T[i]
    }
}

func zigzag(T tabInt, n int) bool {
    if n <= 1 {
        return true
    }

    if T[0] == T[1] {
        return false
    }

    naik := T[0] < T[1]

    for i := 1; i < n-1; i++ {
        if naik && T[i] <= T[i+1] {
            return false
        }

        if !naik && T[i] >= T[i+1] {
            return false
        }

        naik = !naik
    }

    return true
}

func main() {
    var T tabInt
    var n int

    fmt.Scan(&n)
    isiData(&T, n)

    fmt.Println("Array awal:")
    for i := 0; i < n; i++ {
        fmt.Print(T[i], " ")
    }

    if zigzag(T, n) {
        fmt.Println("\nZigzag")
    } else {
        fmt.Println("\nBukan zigzag")
    }

    balik(&T, n)

    fmt.Println("Array setelah dibalik:")
    for i := 0; i < n; i++ {
        fmt.Print(T[i], " ")
    }
    fmt.Println()
}