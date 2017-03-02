package main

import "fmt"

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprint("[%s] %s %s", check, t.Title, t.Deadline)
}
func main() {

}
