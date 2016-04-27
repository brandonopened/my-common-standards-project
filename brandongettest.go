package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"net/http"
)

var data = [][]string{{"Line1", "Grade"}, {"Line2", "Resource"}}

func main() {
	res, err := http.Get("http://api.commonstandardsproject.com/api/v1/standard_sets/49FCDFBD2CF04033A9C347BFA0584DF0_D2604890_grade-01?api-key=pkGLztAUWgJHZv5YbpHoDsV6")
	if err != nil {
		log.Fatal(err)
         
        file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)

    for _, value := range data {
        err := writer.Write(value)
        checkError("Cannot write to file", err) 

}
