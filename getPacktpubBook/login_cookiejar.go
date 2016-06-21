package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

//
// func fetchHiddenLoginFormInputs(body string) {
//
// }
//
// func loginAndFetchPage() (string, error) {
// 	resp, err := http.PostForm(packtPubFreeURL,
// 		url.Values{
// 			"email":    {os.Getenv("PACKTPUB_EMAIL")},
// 			"password": {os.Getenv("PACKTPUB_PASSWORD")},
// 		})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != 200 {
// 		return "", errors.New("Login did not return HTTP 200. Aborting.")
// 	}
//
// 	cookies = resp.Cookies()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return string(body), nil
// }
//
// func claimBook() {
// 	u, _ := url.Parse(packtPubBaseURL)
// 	cookieJar, _ := cookiejar.New(nil)
// 	cookieJar.SetCookies(u, cookies)
//
// 	fmt.Printf("%v\n", cookieJar)
//
// 	client := &http.Client{
// 		Jar: cookieJar,
// 	}
// 	resp, err := client.Get(packtPubBaseURL + freeBookURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
//
// 	fmt.Println(packtPubBaseURL + freeBookURL)
//
// 	if resp.StatusCode >= 400 {
// 		e := errors.New("Claiming book resulted in status >= HTTP 400 response. Aborting")
// 		fmt.Println("URL: ", packtPubBaseURL+freeBookURL)
// 		fmt.Println("Status: ", resp.Status)
// 		log.Fatal(e)
// 	}
// }
// func main() {
// 	claimBook()
// }
var cookies []*http.Cookie

func main() {
	casurl, _ := url.Parse(packtPubBaseURL)
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(casurl, cookies)

	var data = url.Values{}
	data.Add("op", "Login")
	data.Add("email", "forteleaf@gmail.com")
	data.Add("password", "shuria40")

	client := &http.Client{
		Jar: cookieJar,
	}
	// Login
	req, err := http.NewRequest("POST", "packtPubBaseURL", nil)
	req.SetBasicAuth("forteleaf@gmail.com", "shuria40")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	cookies := &http.Clinet{Jar: jar}
	fmt.Println(jar.Cookies(casurl))
	fmt.Println(cookies.jar)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	// // http.Header.Add("User-Agent", "Mozilla/5.0")
	// fmt.Println(resp.Status)
	// // res, err := resp.Get(packtPubBaseURL, "/freelearning-claim/17415/21478")
}
