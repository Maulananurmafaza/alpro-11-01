package main

import "fmt"
func main(){
	var umur int8
	var suhu float64

	suhu = 36.3
	umur = 10
	fmt.Scan(&suhu)
	fmt.Scan(&umur)

	fmt.Println("umur: ", umur)
	fmt.Println("suhu: ", suhu)
	fmt.Println("Alamat memori dari Var suhu ", &suhu)
	fmt.Println("Alamat memori dari Var umur ", &umur)


}