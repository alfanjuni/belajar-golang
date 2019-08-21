package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	var decimalNumber = 2.62
	var exist bool = true

	var message string = "Halo"
	var message2 string = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Printf("bilangan positif : %d \n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Printf("bialngan desimal ; %f\n", decimalNumber)
	fmt.Printf("bilangan desimal : %.3f\n", decimalNumber)
	fmt.Printf("exist ? %t \n", exist)
	fmt.Printf("message: %s \n", message)
	fmt.Println(message2)
}
