package main 

import "fmt"

const NMAX int = 999
type gempa struct {
	nama string
	tahun int
	magnitudo float64
}
type tabGempa [NMAX] gempa

func kategori (m float64)string {
	if m < 4.0 {
		return "Minor"
	} else if m >= 4.0 && m < 6.0 {
		return "Sedang"
	} else if m >= 6.0 && m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}
func rataMagnitudo(T tabGempa, n int) float64 {
    jumlah := 0.0

    for i := 0; i < n; i++ {
        jumlah += T[i].magnitudo
    }
    return jumlah / float64(n)
}
func idxTerkuat (T tabGempa, n int)int {
	idx := 0

	for i := 1; i < n; i++ {
        if T[i].magnitudo > T[idx].magnitudo {
            idx = i
        }
    }
    return idx
}
func hitungKategori (T tabGempa, n int, k string) int {
	jumlah := 0

	for i := 0; i < n; i++ {
        if kategori(T[i].magnitudo) == k {
            jumlah++
        }
    }
    return jumlah
}
func main(){
	var T tabGempa 
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}
	for i := 0; i < n; i++ {
        fmt.Printf("%s %d %.1f %s\n",
            T[i].nama,
            T[i].tahun,
            T[i].magnitudo,
            kategori(T[i].magnitudo))
    }
	rata := rataMagnitudo(T, n)
    fmt.Printf("Rata-rata magnitudo: %.2f\n", rata)

	idx := idxTerkuat(T, n)

    fmt.Printf("Gempa terkuat: %s %d %.1f\n",
        T[idx].nama,
        T[idx].tahun,
        T[idx].magnitudo)

	besar := hitungKategori(T, n, "Besar")
    fmt.Printf("Jumlah gempa kategori Besar: %d\n", besar)

}
