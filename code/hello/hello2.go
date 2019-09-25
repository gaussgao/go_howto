/*
	$ go run hello2.go
*/

package main

import "fmt"
import "./myMath"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(mathClass.Add(1,1))
	fmt.Println(mathClass.Sub(1,2))
}
