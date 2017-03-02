package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getJsonFile, err := http.Get("http://api.geonames.org/weatherJSON?north=44.1&south=-9.9&east=-22.4&west=55.2&username=demo")
	if err != nil {
		log.Println(err)
	}
	defer getJsonFile.Body.Close()

	content, err := ioutil.ReadAll(getJsonFile.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", string(content))
}
