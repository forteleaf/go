package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "host:port")
        os.Exit(1)
    }

    url := os.Args[1]

    response, err := http.Get(url)
    checkError(err)

    if response.StatusCode != 200 {
        fmt.Println(response.Status)
        os.Exit(1)
    }
    b, _ := httputil.DumpResponse(response, false)
    fmt.Printf(string(b))

    var buf [1024]byte
    reader := response.Body
    for {
        n, err := reader.Read(buf[0:])
        checkError(err)
        fmt.Print(string(buf[0:n]))
    }
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
        os.Exit(1)
    }
}
