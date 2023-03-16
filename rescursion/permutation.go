package main

import "fmt"

//Testing code
func main() {
    var data [3]int
    for i := 0; i < 3; i++ {
        data[i] = i
      }
	  getPermutation(data[:], 0, 3)
}

func getPermutation(data []int, i int, length int){
	if length == i {
		fmt.Printf("%v\n",data)
		return
	}

	for j:=i; j < length; j++ {
		data[i], data[j] = data[j], data[i]
		getPermutation(data, i + 1, length)
		data[i], data[j] = data[j], data[i]
	}
}