package main

import "fmt"

const NMAX int = 100
type book struct {
	code string
	title string
	borrowed int
	pages int
}

type books [NMAX]book

func bacaData (T *books, n *int) {
	*n = 0
	fmt.Scan(&T[*n].code)

	for T[*n].code != "SELESAI" {
		fmt.Scan(&T[*n].title,
			&T[*n].borrowed,
			&T[*n].pages)
		*n++
		fmt.Scan(&T[*n].code)
	}
}

func cetakData (T books, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].code,
			T[i].title,
			T[i].borrowed,
			T[i].pages)
	}
}

func hapusBuku (T *books, n *int, X int) {
	i := 0

	for i < *n {
		if T[i].borrowed < X {
			for j := i; j < *n-1; j++ {
				T[j] = T[j+1]
			}
			*n--
		} else {
			i++
		}
	}
}

func totalPinjam (T books, n int) int {
	total := 0

	for i := 0; i < n; i++ {
		total += T[i].borrowed
	}

	return total
}

func idxFavorit(T books, n int) int {
	idx := 0

	for i := 1; i < n; i++ {
		if T[i].borrowed > T[idx].borrowed ||
			(T[i].borrowed == T[idx].borrowed &&
				T[i].pages < T[idx].pages) {
			idx = i
		}
	}

	return idx
}

func main() {
    var T books
    var n, X int

    bacaData(&T, &n)

    fmt.Println("Data awal buku:")
    cetakData(T, n)

    fmt.Scan(&X)

    hapusBuku(&T, &n, X)

    fmt.Println("Data setelah penghapusan:")

    if n == 0 {
        fmt.Println("Tidak ada buku tersisa")
        return
    }

    cetakData(T, n)

    fmt.Println("Total peminjaman:",
        totalPinjam(T, n))

    idx := idxFavorit(T, n)

    fmt.Println("Buku paling sering dipinjam:")
	fmt.Println(T[idx].code,
    T[idx].title,
    T[idx].borrowed,
    T[idx].pages)
}