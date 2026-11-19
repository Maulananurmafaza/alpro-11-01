package main

import "fmt"

func kemasDonat(total int, p12, p6, p1 *int) {
	*p12 = total / 12
	total = total % 12
	*p6 = total / 6
	total = total % 6
	*p1 = total
}
func main() {
	var total, p12, p6, p1 int

	fmt.Scan(&total)

	kemasDonat(total, &p12, &p6, &p1)

	fmt.Println(p12, "paket isi 12")
	fmt.Println(p6, "paket isi 6")
	fmt.Println(p1, "satuan")
}
