package strCat

import "fmt"

func Example_strCat() {
	s := "abc"
	ps := &s // 주소값을
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}
