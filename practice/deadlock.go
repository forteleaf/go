// deadlock(교착상태)와 race condition(경쟁상태) 해결
// http://blog.remotty.com/blog/2015/08/15/golangeuro-anjeonhan-seobiseu-mandeulgi/

package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflictiong access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
