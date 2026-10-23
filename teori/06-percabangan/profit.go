package main

import"fmt"

func main(){
	var bulanSebelumnya, bulanIni float64
	fmt.Scan (&bulanSebelumnya, &bulanIni)

	if bulanIni > bulanSebelumnya {
		fmt.Println("Peningkatan sebesar", bulanIni - bulanSebelumnya)
	} else if bulanIni < bulanSebelumnya {
		fmt.Println("Penurunan sebesar", bulanSebelumnya - bulanIni)
	} else {
		fmt.Println("Tetap")
	}
}