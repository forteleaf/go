// getMacAddress
package main

import (
	"fmt"
	"net"
	"strings"
)

// get MacAddr
func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("error")
	}

	for _, v := range interfaces {
		fmt.Println(v)
		if v.Name == "en0" {
			fmt.Println(v.HardwareAddr)
			ss := strings.Split(v.HardwareAddr.String(), ":")
			fmt.Println(strings.Join(ss, ""))
		}
		// fmt.Println(v.HardwareAddr)
		// fmt.Println(v.Name, v.HardwareAddr)
	}

	/*
			Result
		   0 {1 16384 lo0  up|loopback|multicast}
		   1 {2 1280 gif0  pointtopoint|multicast}
		   2 {3 1280 stf0  0}
		   3 {4 1500 en0 68:a8:6d:0b:08:82 up|broadcast|multicast}
		   4 {5 1500 en1 b2:00:19:3b:23:80 up|broadcast}
		   5 {6 2304 p2p0 0a:a8:6d:0b:08:82 up|broadcast|multicast}
		   6 {7 1500 bridge0 6a:a8:6d:b0:02:00 up|broadcast|multicast}
	*/

}
