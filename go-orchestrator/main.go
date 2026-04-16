package main

import (
	"fmt"
	splitter "orchestrator/internal/splitter"
)



func main() {
	content := "Slices are an important data type in Go, giving a more powerful interface to sequences than arrays. To create a slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make."
	model := "gpt-4"
	

	result, err := splitter.TokenChunker(content, 10, 3, model)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("length of the result: ", len(result))

}
