package main

import "fmt"

type geoCircle struct {
	area, circumference float64
}

type circle struct {
	radius float64
	color string
	property geoCircle
}

func isiData (c *circle) {
	fmt.Scan(&c.radius, &c.color)
}

func hitung(c *circle) {
    c.property.area = 3.14 * c.radius * c.radius
    c.property.circumference = 2 * 3.14 * c.radius
}

func cetak(c circle) {
    fmt.Printf("%.2f %.2f\n",
        c.property.area,
        c.property.circumference)
}

func main() {
    var c circle

    isiData(&c)
    hitung(&c)
    cetak(c)
}