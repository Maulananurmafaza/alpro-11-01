package main

import "fmt"

func add(a, b int32) int32 { // fungsi punya keyward tipe data kembali dan return
	return a + b
}

func eating() { // prosedur punya keyward dan tidak return
	fmt.Println("Saya lapar")
	fmt.Println("Saya ngambil piring")
	fmt.Println("Saya ambil nasi dan beberapa lauk")
	fmt.Println("Saya makan nasi dan lauk")
}

func hitungF(x, y int) {
	var hasil int
	hasil = 2+x + y - 5
	fmt.Println(hasil)
} // prosedur

func hitungFF(x, y int) int {
	var hasil int
	hasil = 2+x + y - 5
	return hasil
} // fungsi

func main() {
	fmt.Println(hitungFF(2, 3))
	x := 2
	y := 3
	hitungF(2, 3)
	hitungF(x, y)
	hitungF(x, 3)
	fmt.Println(add(2, 3)) // panggil fungsi
	eating() // panggil prosedur
}