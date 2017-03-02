package main

import (
	"io/ioutil"
	"net/http"
)

func slurp(url string) (string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	s, err := slurp(url)

}
