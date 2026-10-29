package main

import "fmt"

func main() {
    var tanggal, bulan, tahun int
    fmt.Scan(&tanggal, &bulan, &tahun)

    kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)

    if bulan < 1 || bulan > 12 {
        fmt.Println("Tanggal tidak valid")
        return
    }

    hari := 31

    if bulan == 2 {
        hari = 28
        if kabisat {
            hari = 29
        }
    } else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
        hari = 30
    }
    if tanggal < 1 || tanggal > hari {
        fmt.Println("Tanggal tidak valid")
        return
    }
    fmt.Println("Valid")
    tanggal++

    if tanggal > hari {
        tanggal = 1
        bulan++

        if bulan > 12 {
            bulan = 1
            tahun++
        }
    }

    fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}
