package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		condition(i)
	}
}
func condition(i int) {
	if i%2 == 0 {
		println("-偶数")
	} else {
		println("-奇数")
	}
}
