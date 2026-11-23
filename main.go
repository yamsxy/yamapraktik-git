package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu Celsius: ")
	_, err := fmt.Scan(&celsius)

	if err != nil {
		fmt.Println("Input harus berupa angka!")
		return
	}

	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	fmt.Println("Hasil Konversi:")
	fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Kelvin: %.2f\n", kelvin)
}
