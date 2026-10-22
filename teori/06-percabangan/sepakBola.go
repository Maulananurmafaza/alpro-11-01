package main

import "fmt"

func main() {
	var gol1, gol2, gol3, gol4 int

	fmt.Print("Masukkan jumlah gol 4 tim (pisahkan dengan spasi): ")
	_, err := fmt.Scan(&gol1, &gol2, &gol3, &gol4)
	if err != nil {
		fmt.Println("Masukan tidak valid.")
		return
	}
	terbanyak := gol1
	if gol2 > terbanyak {
		terbanyak = gol2
	}
	if gol3 > terbanyak {
		terbanyak = gol3
	}
	if gol4 > terbanyak {
		terbanyak = gol4
	}

	tersedikit := gol1
	if gol2 < tersedikit {
		tersedikit = gol2
	}
	if gol3 < tersedikit {
		tersedikit = gol3
	}
	if gol4 < tersedikit {
		tersedikit = gol4
	}

	fmt.Printf("Jumlah gol terbanyak: %d\n", terbanyak)
	fmt.Printf("Jumlah gol tersedikit: %d\n", tersedikit)
}