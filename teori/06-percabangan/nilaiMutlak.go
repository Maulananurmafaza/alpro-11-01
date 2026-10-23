package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}
	fmt.Println(n)
}