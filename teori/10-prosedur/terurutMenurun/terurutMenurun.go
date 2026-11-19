package main

import "fmt"

func urutMenurun(a, b *int) {
    if *a < *b {
        temp := *a
        *a = *b
        *b = temp
    }
}
func main() {
    var a, b int

    fmt.Scan(&a, &b)

    for a != b {
        urutMenurun(&a, &b)
        fmt.Println(a, b)

        fmt.Scan(&a, &b)
    }
}
