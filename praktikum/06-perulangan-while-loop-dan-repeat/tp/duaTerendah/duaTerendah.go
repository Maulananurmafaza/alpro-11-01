package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	siswa1, siswa2 := "", ""
	nilai1, nilai2 :=  999, 999

	for i := 0; i < n; i++ {
		var siswa string
		var nilai int
		fmt.Scan(&siswa, &nilai)

		if nilai < nilai1 {
			nilai2 = nilai1
			siswa2 = siswa1
			nilai1 = nilai
			siswa1 =  siswa
		} else if nilai < nilai2 {
			nilai2 = nilai
			siswa2 = siswa
		}
	}
	fmt.Println(siswa1)
	fmt.Println(siswa2)
}
