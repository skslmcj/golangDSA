package main

import "fmt"

func main(){
	fmt.Println(findFib(10))
}

func findFib(i int) int {
	if i < 1 {
		return 0
	}

	if i < 2 {
		return 1
	}

	return findFib(i - 1) + findFib(i - 2)
}