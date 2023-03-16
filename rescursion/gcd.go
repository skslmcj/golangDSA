package main

import "fmt"

func main(){
	fmt.Println(findGcd(24, 18))
}

func findGcd(m int, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	return findGcd(n, m % n)
}