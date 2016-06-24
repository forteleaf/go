package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	URLBase = "https://www.packtpub.com/"
	URLFree = "https://www.packtpub.com/packt/offers/free-learning"
)

// GetURL get Free Book URL
func GetURL() (retUrl string, retErr error) {
	doc, err := goquery.NewDocument(URLFree)
	if err != nil {
		return "", err
	}
	doc.Find("a.twelve-days-claim").Each(func(i int, s *goquery.Selection) {
		node := s.First()
		retUrl, _ = node.Attr("href")
	})
	return retUrl, nil
}

func main() {
	url, _ := GetURL()
	fmt.Println(URLBase + url)

	// // Parse + String preserve the original encoding.
	// resp, err := http.Get("http://www.packtpub.com/packt/offers/free-learning")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // bytes, err := ioutil.ReadAll(resp.Body)
	// // fmt.Println(string(bytes))
	// defer resp.Body.Close()
	// links := All(resp.Body)
	// for _, link := range links {
	// 	fmt.Println(link)
	// }
}

/*
func main() {
	// Sending a literal '%' in an HTTP request's Path
	req := &http.Request{
		Method: "GET",
		Host:   "www.packtpub.com/packt/offers/free-learning", // takes precedence over URL.Host
		URL: &url.URL{
			Host:   "ignored",
			Scheme: "https",
			Opaque: "/%2f/",
		},
		Header: http.Header{
			"User-Agent": {"godoc-example/0.1"},
		},
	}
	out, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Replace(string(out), "\r", "", -1))
}
*/
