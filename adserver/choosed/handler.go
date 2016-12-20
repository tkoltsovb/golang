package main

import (
	"fmt"
	"net/http"
	"io"
	//"io/ioutil"
)

func infoHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello, world!\n")

	fmt.Println("remoteaddr = ", request.RemoteAddr)
	header := request.Header.Get("X-Forwarded-For")
	//if len(header) != 0 {
		fmt.Println("X-Forwarded-For = ", header)
	//}
}