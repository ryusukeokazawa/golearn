package main

func swap2(np, mp *int) {
	*np = 100
	println(np, mp)
	return
}

func swap3(n, m int) {
	println(n, m)
	return
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	swap3(n, m)
	println(n, m) // 関数外(swap2以外)でも反映されている=破壊的操作
}
