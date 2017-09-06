package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	// _ : blank identifier, "I don't want to use this variable"
	res, _ := http.Get("http://www.vecka.nu/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
