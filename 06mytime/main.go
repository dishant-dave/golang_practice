package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("mytime")
	presenttime := time.Now()
	fmt.Println("Present time", presenttime)

	fmt.Println(presenttime.Format("01-02-2006"))
}
