package main
import "fmt"

func main() {
	var a,b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a, b=b, a

	//output
	fmt.Println(a)
	fmt.Println(b)
}