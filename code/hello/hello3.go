// hello3.go
// go run hello3.go
package main

import (
	"fmt"
	"bufio"
	"os"
	
)

func main(){

	inputReader := bufio.NewReader(os.Stdin)
	
	
	for true {
			
		fmt.Println("Please tell me your name :")
		
		input , err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error ocurring :  %s \n",err)
			break
		}	else {
		
			// windows commander should be 2 
			crcfLen := 2;
			// linux or MINGW64 bash should be 1
			crcfLen = 1
			inStr := input[:len(input) - crcfLen]
			
			switch inStr {
			case "Exit":
				fmt.Printf("Byebye\n\n")
				return
			default:
				fmt.Printf("Hello , |%s|\n",inStr)
			}
		}
	}
}