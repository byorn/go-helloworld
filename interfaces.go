package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func testFunc() {
	resp, err := http.Get("http://www.byorns-playground.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

type logWriter struct{}

// * this function implements the Writer interface
// which is why i can pass it into the io.Copy function as a parameter
func (lg logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("wrote bytes", len(bs))
	return len(bs), nil
}
