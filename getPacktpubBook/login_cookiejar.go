package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

var cookies []*http.Cookie

func main() {
	var data = url.Values{}
	data.Add("op", "Login")
	data.Add("email", "forteleaf@gmail.com")
	data.Add("password", "shuria40")

	// python2 에서 사용하는 코드
	// cj = cookielib.CookieJar()
	// opener = urllib2.build_opener(urllib2.HTTPCookieProcessor(cj))
	// urllib2.install_opener(opener)
	// // http.Header.Add("User-Agent", "Mozilla/5.0")
	// fmt.Println(resp.Status)
	// // res, err := resp.Get(packtPubBaseURL, "/freelearning-claim/17415/21478")
	// NewRequest isreturn *http.Request, err
	req, err := http.NewRequest("Get", packtPubBaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	defer req.Body.Close()
	// ioutil.ReadAll(io.Reader) is return []byte, err
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	r

}
