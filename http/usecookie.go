// cookie 를 사용하는 방법에 알아봅시다앙
package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
	// read cookie
	cookie, err := req.Cookie("testcookiename")
	if err == nil {
		cookievalue := cookie.Value
		w.Write([]byte("<b>cookie 값은：" + cookievalue + "</b>\n"))
	} else {
		w.Write([]byte("<b>읽기오류 발생：" + err.Error() + "</b>\n"))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>cookie 설정 성공。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
	// MaxAge : -1 로 설정하면 쿠키는 바로 삭제 된다.
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>cookie 삭제 성공。</b>\n"))
}

func main() {

	http.HandleFunc("/", SayHello)
	http.HandleFunc("/readcookie", ReadCookieServer)
	http.HandleFunc("/writecookie", WriteCookieServer)
	http.HandleFunc("/deletecookie", DeleteCookieServer)

	http.ListenAndServe(":9191", nil)
}
