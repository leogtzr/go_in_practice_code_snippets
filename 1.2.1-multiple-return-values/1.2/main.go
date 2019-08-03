package main

import "fmt"

// Names ...
func Names() (first string, second string) { // multiple return values can have names.
	first = "Hola"
	second = "Mundo"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
