package main

import "fmt"

const NMAX = 1001

type tabNilai [NMAX]int

func inputData(T *tabNilai, n *int) {
    *n = 0

    var nilai int
    fmt.Scan(&nilai)

    for nilai >= 0 && *n < NMAX {
        T[*n] = nilai
        *n++

        fmt.Scan(&nilai)
    }
}

func tertinggi(T tabNilai, n int) int {
    maks := T[0]

    for i := 1; i < n; i++ {
        if T[i] > maks {
            maks = T[i]
        }
    }

    return maks
}

func main() {
    var A, B tabNilai
    var nA, nB int

    inputData(&A, &nA)
    inputData(&B, &nB)

    fmt.Println(tertinggi(A, nA))
    fmt.Println(tertinggi(B, nB))
}