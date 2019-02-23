package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// 根据提示　定义一个初始值　来套用公式
	z := 1.00
	//  临时变量 记录z 上次的值
	temp := 0.00

	for {
		// 计算出最新的z值
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
		if math.Abs(z-temp) < 0.000000000000001 {
			//  当值停止改变（或改变非常小）的时候退出循环
			break
		} else {
			//　赋值最后的值
			temp = z
		}
	}
	return z
}
func main() {
	fmt.Println("牛顿法：", Sqrt(2))
	fmt.Println("math.Sqrt(２):", math.Sqrt(2))
}
