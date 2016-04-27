package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.commonstandardsproject.com/api/v1/standard_sets/49FCDFBD2CF04033A9C347BFA0584DF0_D2604890_grade-01?api-key=pkGLztAUWgJHZv5YbpHoDsV6")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
		fmt.Printf("%s", page)
}
