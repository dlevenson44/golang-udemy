package main

import (
	"fmt"
	"io"
	"os"
)

// class solution
func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

// my solution
// func main() {
// 	f := os.Args[1]
// 	// os.Args shows us arguments passed via run cmd
// 	fmt.Println(f)

// 	//Pass the file name to the ReadFile() function from
// 	//the ioutil package to get the content of the file.

// 	content, error := os.ReadFile(f)

// 	// Check whether the 'error' is nil or not. If it
// 	//is not nil, then print the error and exit.
// 	if error != nil {

// 		log.Fatal(error)
// 	}

// 	// convert the content into a string
// 	str := string(content)

// 	//Print the string str
// 	fmt.Println(str)

// 	os.Open(str)
// }
