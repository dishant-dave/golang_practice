package main

import "fmt"

func main() {
	fmt.Println("array")
	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "orange"
	fruitlist[3] = "blackberry"
	fmt.Println("fruit list is : ", fruitlist)
	fmt.Println("fruit list is : ", len(fruitlist))

	var veglist = [3]string{"apple", "orange", "blackberry"}
	fmt.Println("veg list : ", veglist)
}
