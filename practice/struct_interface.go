// interface 정의
/*
구조체(struct)와 매우 유사한 구조를 갖고 있다.
*/
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, x2, y1, y2 float64
}

/*
	초기화
	var c Circle
	c := new(Circle)
	c := Circle{x:0,y:0,r:5}
	c := Circle{0,0,5}
*/

/*
	필드
	fmt.Println(x.c, c.y, c.r)
	c.x = 10
	c.y = 5
*/

// 어떤 타입에서 인터페이스를 구현하기 위해 반드시 포함하고 있어야 할 메서드의 목록에 해당한다.

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}

func circleArea(c *Circle) float64 {
	// Go 에서는 인자가 복사되기 때문에 원본 변수는 변경되지 않는다.
	return math.Pi * c.r * c.r
}

// receiver 를 추가 . 연산자를 사용해서 해당 함수를 호출할수 있다.
// 이렇게 변경을 하면 Circle 에 대한 포인터를 전달해야 한다는 사실을 자동으로 알고 있다.
// 이 함수는 Circle 하고만 사용할 수 있다.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

// 인터페이스
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// 필드로 사용할 수 잇따.
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	r := Rectangle{0, 10, 0, 10}
	fmt.Println(r.area())
	fmt.Println(totalArea(&c, &r)) // 인터페이스
}
