func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "달기"}
	f3 := append(f1, f2...)     // 이어붙이기
	f4 := append(f1[:2], f2...) // 토마토를 제외하고 이어붙이기
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
}
