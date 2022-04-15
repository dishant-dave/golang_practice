package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome user"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Entre rating for this: ")
	// coma ok
	input, _ := reader.ReadString('\n')
	fmt.Printf("rating is: %T", input)

}
