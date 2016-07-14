package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("gggasdf")
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		return err
	}
}
