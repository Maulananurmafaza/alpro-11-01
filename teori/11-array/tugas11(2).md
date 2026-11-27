### Tugas Minggu 11
#### Maulana Nur Mafaza
#### 109092600015

##### 1. titikTengah
#### Kode
```go
package main

import "fmt"

type titik struct {
	x, y float64
}

func tengah (p1, p2 titik) titik {
	return titik {
		x: (p1.x + p2.x) / 2,
        y: (p1.y + p2.y) / 2,
	}
}

func main() {
    var p1, p2 titik

    fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)

    hasil := tengah(p1, p2)

    fmt.Println(hasil.x, hasil.y)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/11-array/titikTengah/titikTengah_output.png

##### 2. lingkaran
#### Kode
```go
package main

import "fmt"

type geoCircle struct {
	area, circumference float64
}

type circle struct {
	radius float64
	color string
	property geoCircle
}

func isiData (c *circle) {
	fmt.Scan(&c.radius, &c.color)
}

func hitung(c *circle) {
    c.property.area = 3.14 * c.radius * c.radius
    c.property.circumference = 2 * 3.14 * c.radius
}

func cetak(c circle) {
    fmt.Printf("%.2f %.2f\n",
        c.property.area,
        c.property.circumference)
}

func main() {
    var c circle

    isiData(&c)
    hitung(&c)
    cetak(c)
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/11-array/lingkaran/lingkaran_output.png

##### 3. larik zigzag
#### Pseudocode
ALGORITMA LarikZigzag

DEKLARASI
    T : array[128] of integer
    n, i, temp : integer
    naik : boolean

PROCEDURE IsiData(T, n)
    FOR i ← 0 TO n-1 DO
        input(T[i])
    ENDFOR
ENDPROCEDURE

PROCEDURE Balik(T, n)
    FOR i ← 0 TO n/2-1 DO
        temp ← T[i]
        T[i] ← T[n-1-i]
        T[n-1-i] ← temp
    ENDFOR
ENDPROCEDURE

FUNCTION Zigzag(T, n) : boolean
    IF n <= 1 THEN
        RETURN true
    ENDIF

    IF T[0] = T[1] THEN
        RETURN false
    ENDIF

    naik ← T[0] < T[1]

    FOR i ← 1 TO n-2 DO
        IF naik AND T[i] <= T[i+1] THEN
            RETURN false
        ENDIF

        IF NOT naik AND T[i] >= T[i+1] THEN
            RETURN false
        ENDIF

        naik ← NOT naik
    ENDFOR

    RETURN true
ENDFUNCTION

ALGORITMA UTAMA
    input(n)

    IsiData(T, n)

    output(T)

    IF Zigzag(T, n) THEN
        output("Zigzag")
    ELSE
        output("Bukan zigzag")
    ENDIF

    Balik(T, n)

    output(T)

SELESAI

##### 4. nilai ujian
#### Kode
```go
package main

import "fmt"

const NMAX = 1001

type tabNilai [NMAX]int

func inputData(T *tabNilai, n *int) {
    *n = 0

    var nilai int
    fmt.Scan(&nilai)

    for nilai >= 0 && *n < NMAX {
        T[*n] = nilai
        *n++

        fmt.Scan(&nilai)
    }
}

func tertinggi(T tabNilai, n int) int {
    maks := T[0]

    for i := 1; i < n; i++ {
        if T[i] > maks {
            maks = T[i]
        }
    }

    return maks
}

func main() {
    var A, B tabNilai
    var nA, nB int

    inputData(&A, &nA)
    inputData(&B, &nB)

    fmt.Println(tertinggi(A, nA))
    fmt.Println(tertinggi(B, nB))
}
```


##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/11-array/nilaiUjian/nilaiUjian_output.png


##### 5. buku1
#### Kode
```go
package main

import "fmt"

const NMAX = 500

type buku struct {
    kode     string
    judul    string
    penerbit string
    tahun    int
}

type tabBuku [NMAX]buku

func tambahBuku(T *tabBuku, n *int, b buku) {
    if *n < NMAX {
        T[*n] = b
        *n = *n + 1
    } else {
        fmt.Println("Data penuh")
    }
}

func cetak(T tabBuku, n int, x string) {
    for i := 0; i < n; i++ {
        if T[i].penerbit == x {
            fmt.Printf("Kode: %s, Judul: %s, Penerbit: %s, Tahun: %d\n",
                T[i].kode, T[i].judul, T[i].penerbit, T[i].tahun)
        }
    }
}
```

##### buku2
#### Kode
```go
func main() {
    var T tabBuku
    n := 0

    tambahBuku(&T, &n, buku{"BK01", "Ali Bercerita", "Erlangga", 2015})
    tambahBuku(&T, &n, buku{"BK02", "Bumi Berputar", "Gramedia", 2018})
    tambahBuku(&T, &n, buku{"BK03", "Cinta Semesta", "Mizan", 2016})
    tambahBuku(&T, &n, buku{"BK04", "Dunia Kita", "Gramedia", 2020})
    tambahBuku(&T, &n, buku{"BK05", "Elang Muda", "Bentang", 2012})
    tambahBuku(&T, &n, buku{"BK06", "Fajar Baru", "Erlangga", 2019})
    tambahBuku(&T, &n, buku{"BK07", "Gita Cinta", "Gramedia", 2021})
    tambahBuku(&T, &n, buku{"BK08", "Hutan Sunyi", "Mizan", 2014})

    cetak(T, n, "Gramedia")
}
```

##### Output
https://github.com/Maulananurmafaza/alpro-11-01/blob/main/teori/11-array/buku_2/buku_output.png