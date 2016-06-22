package packtFree

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/xmlpath.v2"

	"golang.org/x/net/html"

	"github.com/JSila/packtFree/info"
	"github.com/JSila/packtFree/log"
)

const (
	// URL of packt publishing free offer site
	URL string = "https://www.packtpub.com/packt/offers/free-learning"
	// FILE is filepath where data is located
	FILE string = "book.json"
)

// Data is info about this days free offer
var Data = new(info.Info)

func normalizeHTML(doc []byte) *bytes.Reader {
	reader := bytes.NewReader(doc)
	root, err := html.Parse(reader)

	if err != nil {
		log.Log.Panic("normalizeHTML: can't parse html")
	}

	var b bytes.Buffer
	html.Render(&b, root)
	return bytes.NewReader(b.Bytes())
}

// Scrape pulls data from packt publishing free offer page
func Scrape() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	root, err := xmlpath.ParseHTML(normalizeHTML(body))
	if err != nil {
		return err
	}
	path := xmlpath.MustCompile("//div[contains(@class, 'dotd-main-book-image float-left')]")
	if value, ok := path.String(root); ok {
		img := strings.TrimSpace(value)
		re, _ := regexp.Compile("src=\"(.[^\"]*)\"")
		match := re.FindStringSubmatch(img)
		Data.Img = match[1]
	}

	path = xmlpath.MustCompile("//div[contains(@class, 'dotd-title')]/h2")
	if value, ok := path.String(root); ok {
		Data.Title = strings.TrimSpace(value)
	}

	path = xmlpath.MustCompile("//div[contains(@class, 'dotd-main-book-summary')]/div[not(@class)]")
	iter := path.Iter(root)

	Data.Description = []string{}
	for iter.Next() {
		Data.Description = append(
			Data.Description,
			strings.TrimSpace(iter.Node().String()),
		)
	}

	path = xmlpath.MustCompile("//span[contains(@class, 'packt-js-countdown')]/@data-countdown-to")
	if value, ok := path.String(root); ok {
		t, _ := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
		Data.EndTime = time.Unix(t, 0)
	}

	return nil
}

// IsOutdated determines if saved data is outdated and need to be updated
func IsOutdated() (need bool) {
	file, err := ioutil.ReadFile(FILE)
	if err != nil {
		need = true
	}
	if err = json.Unmarshal(file, &Data); err != nil {
		need = true
	}
	if time.Now().After(Data.EndTime) {
		need = true
	}
	return
}

// UpdateFile saves the data back to the file
func UpdateFile() error {
	jdata, err := json.MarshalIndent(Data, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FILE, jdata, os.ModePerm)
}

// Get gets the data from filesystem or internet if needed
func Get() error {
	if !IsOutdated() {
		return nil
	}
	if err := Scrape(); err != nil {
		return err
	}
	if err := UpdateFile(); err != nil {
		return err
	}
	return nil
}
