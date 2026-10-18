package main

import "fmt"

func main(){
	var nilaiUtama, nilaiRemedial int
	fmt.Scan(&nilaiUtama, &nilaiRemedial)

	if nilaiUtama >= 70 {
		if nilaiUtama >= 85 {
			fmt.Println ("lulus, sangat baik")
		} else {
			fmt.Println("lulus, baik")
		}
	} else {
		if nilaiRemedial >= 70 {
			fmt.Println ("remedial, lulus remedial")
		} else {
			fmt.Println("remedial, tidak lulus")
		}
	}
}
	
	
		
	


