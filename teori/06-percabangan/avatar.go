package main

import "fmt"

func main() {
    var p int
    fmt.Scan(&p)

    appaDewasa := p / 5
    if appaDewasa > 3 {
        appaDewasa = 3
    }

    sisa := p - appaDewasa*5

    if sisa > 0 && appaDewasa < 3 {
        appaDewasa++
        sisa = p - appaDewasa*5
        if sisa < 0 {
            sisa = 0
        }
    }

    appaKecil := sisa / 2
    if appaKecil > 5 {
        appaKecil = 5
    }

    sisa = sisa - appaKecil*2

    fmt.Print("Dewasa ", appaDewasa)

    if appaKecil > 0 {
        fmt.Print(", kecil ", appaKecil)
    }

    if sisa > 0 {
        fmt.Print(", dan ", sisa, " tak berangkat")
    }

    fmt.Println()
}
