package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

// custom type with nothing in it
// it is used as a receive to implement our own io.Write function
// just by doing that, we can say logWriter implements io.Write
type logWriter struct {}

func main(){
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//fmt.Println("Response is: ", resp)

	//// allocate space for the body byte slice
	//// the io.Read will not auto expand the slice, it will just stop writing once its full
	//body := make([]byte, 999999)
	//// populate the body from the reader
	//resp.Body.Read(body)
	//fmt.Println("Body is: ", string(body))

	lw := logWriter{}
	// a better way
	io.Copy(lw, resp.Body)
}

// our own implementation of Write
func (logWriter) Write(bs []byte) (n int, err error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
