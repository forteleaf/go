package main

import (
	"fmt"
	"time"
)

type Deadline time.Time

func (d *Deadline) OverDue() bool {
	// pointer로 바꾸고 난 뒤에
	// return time.Time(d).Before(time.Now())
	return d != nil && time.Time(*d).Before(time.Now())
}

type status int

const (
	UNKNOW status = iota
	TODO
	DONE
)

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}

// func ExampleDeadline_OverDue() {
func main() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"no due", TODO, nil}
	// fmt.Println(d1.OverDue())
	// fmt.Println(d2.OverDue())
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
}
