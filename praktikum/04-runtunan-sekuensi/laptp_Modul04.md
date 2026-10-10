# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtutan dan Sekuensi</h1>
<p align="center">Maulana Nur Mafaza - 109092600015</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

```go
package main

import "fmt"

func main(){
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if intNum > 0 {
		fmt.Println(intNum > 0)
	}
	if intNum >= 5 && intOther < 11 {
		fmt.Println(intNum >= 5 && intOther < 11)
	}
	if sngNum > 0 || (intNum > 0 && -1 * intOther == -10) {
		fmt.Println(sngNum > 0 || (intNum > 0 && -1 * intOther == -10))
	}
}
```

##### Output




#### Deskripsi


### 2. Tracing: Evaluasi Pernyataan Kondisi

```go
package main

import "fmt"

func main(){
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
}

if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output



#### Deskripsi


## Kesimpulan
