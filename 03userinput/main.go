package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter input:");
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for rating: ", input);
	fmt.Printf("Type of this rating is %T", input);

}