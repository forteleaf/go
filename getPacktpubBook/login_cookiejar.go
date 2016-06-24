package main

import "net/http"

const (
	packtPubBaseURL string = "https://www.packtpub.com"
	packtPubFreeURL string = "https://www.packtpub.com/packt/offers/free-learning"
)

func main() {
	cj := http.CookieJar.SetCookies(u, cookies)
}
