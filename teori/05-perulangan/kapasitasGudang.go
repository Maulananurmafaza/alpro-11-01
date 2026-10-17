package main

import "fmt"

func main(){
	var k, v int
	fmt.Scan(&k)
	totalBarang := 0

	for i:= 0; i < 5; i++ {
		fmt.Scan(&v)
		totalBarang += v

		fmt.Println(totalBarang >= k)

	}
}