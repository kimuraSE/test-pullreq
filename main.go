package main

import "fmt"

func main() {
	var p *int
	*p = 1
	fmt.Println(*p)

	if true {
		fmt.Println("true")
	} 
	
}