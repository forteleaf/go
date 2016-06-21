package printBytes
import "fmt"
func Example_printBytes(){
	s:= "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printfBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// eab080eb8298eb8ba4
	// ea b0 89 eb 82 98 eb 8b a4

}