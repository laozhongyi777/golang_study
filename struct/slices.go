package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 从　０开始计算，　打印时忽略最后一个数子
	fmt.Println(s)
}
