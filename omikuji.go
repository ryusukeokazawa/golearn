package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := randomInt(1, 7)
	switch s {
	case 6:
		fmt.Println(strconv.Itoa(s) + "大吉")
	case 5, 4:
		fmt.Println(strconv.Itoa(s) + "中吉")
	case 3, 2:
		fmt.Println(strconv.Itoa(s) + "吉")
	case 1:
		fmt.Println(strconv.Itoa(s) + "凶")
	}
}
