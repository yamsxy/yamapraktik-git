package main

import "fmt"

func main() {
	N := 5
	jumlah := 0

	for i := 1; i <= N; i++ {
		jumlah += i
		fmt.Printf("%d ", i)
	}

	fmt.Println("yama idiot")

	hitung := N

	for hitung >= 1 {
		fmt.Printf("%d ", hitung)
		hitung--
	}

	fmt.Println()
	fmt.Println("Jumlah:", jumlah)
}
