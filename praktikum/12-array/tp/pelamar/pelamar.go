package main

import "fmt"

const NMAX = 100

type applicant struct {
	id string
	name string
	testGrade float64
	testDuration int
}

func salinLolos (A [NMAX]applicant, n int, batas float64) ([NMAX]applicant, int) {
	var B [NMAX]applicant
	m := 0

	for i := 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[m] = A[i]
			m++
		}
	}
	return B, m
}

func rataDurasi (B [NMAX]applicant, m int) float64 {
	if m == 0 {
		return 0
	}
	jumlah := 0

	for i := 0; i < m; i++ {
		jumlah +=B[i].testDuration
	}
	return float64(jumlah) / float64(m)
}

func main () {
	var A[NMAX]applicant
	var n int
	var batas float64

	for n < NMAX {
        fmt.Scan(&A[n].id)

        if A[n].id == "END" {
            break
        }

        fmt.Scan(&A[n].name, &A[n].testGrade, &A[n].testDuration)
        n++
    }

    fmt.Scan(&batas)

    fmt.Println("Data awal:")
    for i := 0; i < n; i++ {
        fmt.Printf("%s %s %.1f %d\n",
            A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
    }

    B, m := salinLolos(A, n, batas)

    fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)

    for i := 0; i < m; i++ {
        fmt.Printf("%s %s %.1f %d\n",
            B[i].id, B[i].name, B[i].testGrade, B[i].testDuration)
    }

    if m == 0 {
        fmt.Println("Tidak ada pelamar yang lolos")
    } else {
        fmt.Printf("Jumlah pelamar lolos: %d\n", m)
        fmt.Printf("Rata-rata durasi: %.2f menit\n",
            rataDurasi(B, m))
    }
}
