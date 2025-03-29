package main

import (
	"fmt"
	"os"
)

func main(){
	message, err := os.Open("./message.txt")
	if err != nil {
		fmt.Printf("No document found !")
	}

	message.Read()


}