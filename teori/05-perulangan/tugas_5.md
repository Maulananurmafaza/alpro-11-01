### Tugas Minggu 5
#### Maulana Nur Mafaza
#### 10909260015

##### 1. Poin Game
#### Kode

```go
package main

import "fmt"

func main(){
	var poin, total int
	total = 0

	fmt.Scan(&poin)

	for poin != 0 {
		total = total + poin
		fmt.Scan(&poin)
	}
	fmt.Println(total)
}
```
##### output


##### 2. Kuadrat Digit
#### Kode

```go
package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	total := 0

	for x > 0 {
		digit := x % 10
		kuadrat := digit * digit

		fmt.Printf("%d", kuadrat)
	total += kuadrat 
	x /= 10
	}
	fmt.Println()
	fmt.Println(total)
}
```
##### output


##### 3. Donat
#### Kode

```go
package main

import "fmt"

func main(){
	var n, m ,x, y int
	fmt.Scan(&n, &m, &x, &y)
	jumlahDonat := 0

	for n >= x && m >=y{
		n -= x
		m -= y
		jumlahDonat++
	
	}
	fmt.Println("Jumlah donat", jumlahDonat)

}
```
##### output


##### 4. Kapasitas Gudang
#### Kode

```go
package main

import "fmt"

func main(){
	var k, v int
	fmt.Scan(&k)
	totalBarang := 0

	for i:= 0; i < 5; i++ {
		fmt.Scan(&v)
		totalBarang += v

		fmt.Println(totalBarang >= k)

	}
}
```
##### output