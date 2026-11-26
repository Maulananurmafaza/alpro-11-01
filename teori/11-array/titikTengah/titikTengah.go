package main

import "fmt"

type titik struct {
	x, y float64
}

func tengah (p1, p2 titik) titik {
	return titik {
		x: (p1.x + p2.x) / 2,
        y: (p1.y + p2.y) / 2,
	}
}

func main() {
    var p1, p2 titik

    fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)

    hasil := tengah(p1, p2)

    fmt.Println(hasil.x, hasil.y)
}