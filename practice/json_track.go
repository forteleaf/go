package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type tracks struct {
	Toptracks []toptracks_info
}

type toptracks_info struct {
	Track []track_info
	Attr  []attr_info
}

type track_info struct {
	Name       string
	Duration   string
	Listeners  string
	Mbid       string
	Url        string
	Streamable []streamable_info
	Artist     []artist_info
	Attr       []track_attr_info
}

type attr_info struct {
	Country    string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}

type streamable_info struct {
	Text      string
	Fulltrack string
}

type artist_info struct {
	Name string
	Mbid string
	Url  string
}

type track_attr_info struct {
	Rank string
}

func get_content() {
	// json data
	url := "http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&api_key=c1572082105bd40d247836b5c1819623&format=json&country=Netherlands"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data tracks
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)
}

func main() {
	get_content()
}
