package main

import (
	"fmt"

	"github.com/metabloxDID/did"
)

func main() {
	document, err := did.CreateDID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created did: ", document)
}
