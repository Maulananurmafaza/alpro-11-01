# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">[Maulana Nur Mafaza] - [109092600015]</p>

### 1. ganjil

```go
package main

import "fmt"

func main(){
	var n int
	jumlah := 0

	fmt.Println("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil: ")


	for i := 1; i <= n; i+= 2 {
		fmt.Println(i, " ")
		jumlah +=i
	}

	fmt.Println()
	fmt.Println("jumlah =", jumlah)
}
```

##### Output



#### Deskripsi

### 2. Perkalian

```go
package main

import "fmt"

func main(){
	var a, b int
	fmt.Println("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Println("Masukkan batas perkalian b: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}
```

##### Output



#### Deskripsi


## Kesimpulan
