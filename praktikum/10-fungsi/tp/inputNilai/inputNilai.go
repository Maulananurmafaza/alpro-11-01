package main

import "fmt"

type student struct {
	nim int
	nama string
	grade float64
}

func cariNIM(T []student, n, nim int) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}
func rataRata(T []student, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}
func main() {
	var T [100]student
	n := 0

	for {
		var input, nama string
		var nilai float64
		var nim int

		fmt.Scan(&input)

		if input == "STOP" {
			break
		}

		fmt.Sscan(input, &nim)
		fmt.Scan(&nama, &nilai)

		if cariNIM(T[:], n, nim) == -1 {
			T[n] = student{nim, nama, nilai}
			n++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].nama, T[i].grade)
	}
	if n > 0 {
		fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T[:], n))
	}
}

