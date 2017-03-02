package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func JsonT() {
	jsonfile, err := http.Get("http://api.geonames.org/weatherJSON?north=44.1&south=-9.9&east=-22.4&west=55.2&username=demo")
	defer jsonfile.Body.Close()

	content, err := ioutil.ReadAll(jsonfile.Body)
	if err != nil {
		log.Println("err", err)
	}

	var parsed map[string]interface{} // json 문서를 저장할 곳
	json.Unmarshal(content, &parsed)  // content -> map
	// jsonData, err := json.MarshalIndent(parsed, "", "  ") // map -> json
	if err != nil {
		panic(err)
	}
	// for _, value := range parsed {
	// 	fmt.Println(value)
	// }
	fmt.Println(reflect.TypeOf(parsed))
	fmt.Println("os.Args[0]", os.Args[0]+" <json filename>")
	// fmt.Println(string(jsonData))

}

func JsonTx() {
	doc := `
	{
		"name": "forteleaf",
		"age": 99,
		"add": "i don't know"
	}
	`

	var data map[string]interface{}    // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
	json.Unmarshal([]byte(doc), &data) // doc를 바이트 슬라이스로 변환하여 넣고,
	// data의 포인터를 넣어줌

	fmt.Println(data["name"], data["age"], data["add"]) // maria 10: 맵에 키를 지정하여 값을 가져옴
	fmt.Printf("%+v\n", data)
}

func JsonMe() {
	jsonfile, err := http.Get("http://api.geonames.org/weatherJSON?north=44.1&south=-9.9&east=-22.4&west=55.2&username=demo")
	if err != nil {
		log.Println(err)
	}
	defer jsonfile.Body.Close()

	content, err := ioutil.ReadAll(jsonfile.Body)
	if err != nil {
		log.Println("err", err)
	}

	var parsed map[string]interface{} // json 문서를 저장할 곳
	json.Unmarshal(content, &parsed)  // content -> map
	if err != nil {
		panic(err)
	}
	// refelct.TypeOf 타입 확인
	// fmt.Println(reflect.TypeOf(parsed))
	weatherObservations := parsed["weatherObservations"]
	fmt.Println(weatherObservations["ICAO"])

	// Go -> JSON 보기좋게 변환
	// doc, _ := json.MarshalIndent(parsed, "", "  ")
	// fmt.Println(string(doc))

	// map pared 출력
	// for key, value := range parsed {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println("os.Args[0]", os.Args[0]+" <json filename>")
}
func main() {
	JsonMe()
}
