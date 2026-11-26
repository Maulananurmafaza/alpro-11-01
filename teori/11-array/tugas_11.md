### Tugas Minggu 11
#### Maulana Nur Mafaza
#### 109092600015

##### 1. titikTengah
#### Pseudocode
ALGORITMA TitikTengah

DEKLARASI
    p1, p2, hasil : titik

PROCEDURE HitungTengah(p1, p2, hasil)
    hasil.x ← (p1.x + p2.x) / 2
    hasil.y ← (p1.y + p2.y) / 2
ENDPROCEDURE

ALGORITMA UTAMA
    input(p1.x, p1.y, p2.x, p2.y)

    HitungTengah(p1, p2, hasil)

    output(hasil.x, hasil.y)

SELESAI

##### 2. lingkaran
#### Pseudocode
ALGORITMA Lingkaran

DEKLARASI
    TYPE geoCircle :
        area, circumference : real
    ENDTYPE

    TYPE circle :
        radius : real
        color : string
        property : geoCircle
    ENDTYPE

    c : circle

PROCEDURE IsiData(c)
    input(c.radius, c.color)
ENDPROCEDURE

PROCEDURE Hitung(c)
    c.property.area ← 3.14 * c.radius * c.radius
    c.property.circumference ← 2 * 3.14 * c.radius
ENDPROCEDURE

PROCEDURE Cetak(c)
    output(c.property.area, c.property.circumference)
ENDPROCEDURE

ALGORITMA UTAMA
    IsiData(c)
    Hitung(c)
    Cetak(c)

SELESAI

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
#### Pseudocode
ALGORITMA NilaiUjian

DEKLARASI
    CONST NMAX = 1001
    TYPE tabNilai : array[0..NMAX-1] of integer

    T : tabNilai
    n : integer

PROCEDURE inputData(T, n)
    DEKLARASI
        nilai : integer
        i : integer

    ALGORITMA
        n ← 0

        FOR i ← 0 TO NMAX-1 DO
            input(nilai)

            IF nilai < 0 THEN
                RETURN
            ENDIF

            T[n] ← nilai
            n ← n + 1
        ENDFOR

ENDPROCEDURE

ALGORITMA UTAMA
    inputData(T, n)

SELESAI

##### 5. buku1
#### Pseudocode
ALGORITMA Buku1

DEKLARASI
    TYPE buku :
        kode, judul, penerbit : string
        tahun : integer
    ENDTYPE

    TYPE tabBuku : array[1..500] of buku

PROCEDURE cetak(T, n, X)
    {I.S. Terdefinisi array T yang berisi n buku}
    {F.S. Menampilkan buku yang diterbitkan oleh penerbit X}

    FOR i ← 1 TO n DO
        IF T[i].penerbit = X THEN
            output(T[i].kode, T[i].judul,
                   T[i].penerbit, T[i].tahun)
        ENDIF
    ENDFOR
ENDPROCEDURE

SELESAI

##### buku2
#### Pseudocode
ALGORITMA Buku2

DEKLARASI
    T : tabBuku
    n, i : integer
    X : string

ALGORITMA
    n ← 10

    FOR i ← 1 TO n DO
        input(T[i].kode)
        input(T[i].judul)
        input(T[i].penerbit)
        input(T[i].tahun)
    ENDFOR

    X ← "Gramedia"

    cetak(T, n, X)

SELESAI