package main

import "fmt"

const NMAX = 500

type buku struct {
    kode     string
    judul    string
    penerbit string
    tahun    int
}

type tabBuku [NMAX]buku

func tambahBuku(T *tabBuku, n *int, b buku) {
    if *n < NMAX {
        T[*n] = b
        *n = *n + 1
    } else {
        fmt.Println("Data penuh")
    }
}

func cetak(T tabBuku, n int, x string) {
    for i := 0; i < n; i++ {
        if T[i].penerbit == x {
            fmt.Printf("Kode: %s, Judul: %s, Penerbit: %s, Tahun: %d\n",
                T[i].kode, T[i].judul, T[i].penerbit, T[i].tahun)
        }
    }
}

func main() {
    var T tabBuku
    n := 0

    tambahBuku(&T, &n, buku{"BK01", "Ali Bercerita", "Erlangga", 2015})
    tambahBuku(&T, &n, buku{"BK02", "Bumi Berputar", "Gramedia", 2018})
    tambahBuku(&T, &n, buku{"BK03", "Cinta Semesta", "Mizan", 2016})
    tambahBuku(&T, &n, buku{"BK04", "Dunia Kita", "Gramedia", 2020})
    tambahBuku(&T, &n, buku{"BK05", "Elang Muda", "Bentang", 2012})
    tambahBuku(&T, &n, buku{"BK06", "Fajar Baru", "Erlangga", 2019})
    tambahBuku(&T, &n, buku{"BK07", "Gita Cinta", "Gramedia", 2021})
    tambahBuku(&T, &n, buku{"BK08", "Hutan Sunyi", "Mizan", 2014})

    cetak(T, n, "Gramedia")
}