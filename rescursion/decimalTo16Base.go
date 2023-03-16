package main

import "fmt"

func main(){
	convertToHex(95)
	fmt.Println()
}

func convertToHex(i int) {
	hexBase := "0123456789ABCDEF"
	base := 16
	digit := i % base

	i = i / base

	if i != 0 {
		convertToHex(i)
	}
	fmt.Print(string(hexBase[digit]))
}