package main

import "fmt"

func main() {
	p := [...]int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(p); i++ {
		sum = sum + p[i]
	}
	fmt.Println(sum)
}
