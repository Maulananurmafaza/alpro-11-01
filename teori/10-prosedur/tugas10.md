### Tugas Minggu 10
#### Maulana Nur Mafaza
#### 109092600015

##### 1. paketDonat
#### Pseudocode
program PaketDonat
kamus
    total, p12, p6, p1 : integer

algoritma
    input(total)
    kemasDonat(total, p12, p6, p1)

    output(p12, " paket isi 12")
    output(p6, " paket isi 6")
    output(p1, " satuan")

##### 2. konversiBerat
#### Pseudocode
program KonversiBerat
kamus
    kg : real
    gram, ons, pon : integer

algoritma
    input(kg)
    konversiBerat(kg, gram, ons, pon)

    output(gram, "g")
    output(ons, " ons")
    output(pon, " pon")

##### 3. terurutMenurun
#### Pseudocode
procedure mengurutkanMenurun(in/out a, b : integer)
{ I.S. terdefinisi bilangan bulat a dan b }
{ F.S. nilai a > b, pertukaran nilai a dan b mungkin dilakukan }

ALGORITMA
    if a < b then
        temp <- a
        a <- b
        b <- temp
    endif
endprocedure

ALGORITMA UTAMA
    input(a, b)

    while a <> b do
        mengurutkanMenurun(a, b)
        output(a, b)
        input(a, b)
    endwhile

##### 4. Soal Tracing
a. in p, q : integer
Tidak valid, karena parameter in hanya dapat menerima nilai dan tidak boleh diubah.

b. in p : integer, in/out q : integer
Nilai awal p = 5 dan q = 5.
 p = 5 * 2 = 10.
 q = 10 + 5 = 15.
Karena q merupakan parameter in/out, nilai m berubah menjadi 15.
Output = 15

c. in/out p : integer, in q : integer
 Nilai awal p = 5 dan q = 5.
 p = 5 * 2 = 10.
 q = 10 + 5 = 15.
Karena p merupakan parameter in/out, nilai m berubah menjadi 10.
Output = 10

d. in/out p, q : integer
 Nilai awal p = 5 dan q = 5.
 p = 5 * 2 = 10.
 q = 10 + 10 = 20.
Karena p dan q sama-sama mengacu pada variabel m, nilai akhirnya menjadi 20.
Output = 20