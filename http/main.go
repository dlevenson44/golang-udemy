package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := []byte{resp.Body}
	// make function takes slice type and number of elements we want slice to be initiate with
	// we use 99999 because Read function doesn't resize slice if it's full
	// we assume 99999 elements is enough elements for our response
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// provide output func as first arg
	// provide input data as second arg
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote ", len(bs), " bytes.")
	return len(bs), nil
}
