package main

import "fmt"

/*
	var a map[string]int = make (map[string]int)
	var b = make(map[string]int)
	c := make(map[string]int)
*/
func map1() {
	a := map[string]int{"hello": 10, "world": 20}
	b := map[string]int{
		"hello": 10,
		"world": 20,
	}
	fmt.Println(a["hello"])
	fmt.Println(b["world"])

}

func map2() {
	solarSystem := make(map[string]float32) // key = string, value = float32
	solarSystem["Mercury"] = 87.969
	solarSystem["Earth"] = 365.33

	fmt.Println(solarSystem["Mercury"])
	fmt.Println(solarSystem["Mery"]) // 존재하지 않는 값이면 float32의 초기값이 출력

	// 키값이 있는지 없는지 검사
	value, ok := solarSystem["Earth"]

	if value, ok := solarSystem["Earth"]; ok {
		fmt.Println("조건을 검사한 후에", value)
	}

	fmt.Println(value, ok)
	for _, sth := range solarSystem {
		fmt.Println(sth)
	}
}

// map 순회하기
func map3_for() {
	solarSystem := make(map[string]float32)
	solarSystem["Mercury"] = 87.969
	solarSystem["Earth"] = 365.33

	//map 순회하기
	for key, value := range solarSystem {
		fmt.Println(key, value)
	}
}

// map 삭제하기
func map4_delete() {
	a := map[string]int{"Hello": 10, "world": 20}
	delete(a, "world") // map a 에서 world 키 삭제
	fmt.Println(a)
}

// map in map
func main() {
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676E+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219E+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185E+23,
			"orbitalPeriod": 686.9600,
		},
	}
	fmt.Println(terrestrialPlanet["Mars"]["mass"])
	for key, value := range terrestrialPlanet {
		fmt.Println(key, value)
	}
}
