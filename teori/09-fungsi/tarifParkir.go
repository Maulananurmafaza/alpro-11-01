package main

import "fmt"

func tarif(jam int) int{
	if jam <= 2{
		return 5000
	} else {
	return 5000 + (jam - 2) * 3000
	}
}
func main(){
	 var jam int
	 
	 fmt.Println("Lama parkir: ")
	 fmt.Scan(&jam)

	 biaya := tarif(jam)
	 fmt.Println("Total tarif parkir = Rp", biaya)
}