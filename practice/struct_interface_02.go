package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("hello, my name is ", p.Name, "ya~")
}

// Person 객체의 메소드를 포함하고 있는 형태
type Android struct {
	Person
	Model string
}

// 위와 아래는 엄연히 다른 방식 이다.
type Android struct {
	Person Person
	Model  string
}

func main() {
	p := Person{"hohohoh"}
	p.talk()
	// hello, my name is hohohoh, ya~
	a := new(Android)
	a.Person.talk()
	a.talk()
}
