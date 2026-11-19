package main

import "fmt"

const NMAX int = 2002

type tabNIM [NMAX]int

func cariNim(T tabNIM, n, x int, idx *int) {
    *idx = -1

    for i := 0; i < n; i++ {
        if T[i] == x {
            *idx = i
            return
        }
    }
}
func ubahNIM(T *tabNIM, idx, baru int) {
    T[idx] = baru
}
func cetak(T tabNIM, n int) {
    for i := 0; i < n; i++ {
        fmt.Println(T[i])
    }
}
func main() {
    var T tabNIM
    var n, x, idx, baru int

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&T[i])
    }

    fmt.Scan(&x)

    cariNim(T, n, x, &idx)

    if idx != -1 {
        fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)

        fmt.Scan(&baru)

        ubahNIM(&T, idx, baru)

        fmt.Println("Daftar NIM setelah terjadi perubahan data:")
        cetak(T, n)

    } else {
        fmt.Println("NIM tidak ditemukan dalam sistem")
    }
}