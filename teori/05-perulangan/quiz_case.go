package main

import "fmt"

func main() {
	var n, prev, count int
	if fmt.Scan(&n) == 1 {
		prev = n
		count = 1

		for count < 10 {
			var current int
			if fmt.Scan(&current) != 1 {
				break
			}
			count++

			if current > prev {
				break
			}
			prev = current
		}
	}
	fmt.Println(count)

}