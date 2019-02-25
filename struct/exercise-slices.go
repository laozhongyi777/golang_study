package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//外层sliece
	a := make([][]uint8, dy)
	for x := range a {
		//里层slice
		b := make([]uint8, dx)
		for y := range b {
			//给里层ｓｌｉｃｅ 每个元素赋值
			b[y] = uint8(x*y - 1)
		}
		a[x] = b
	}
	return a
}

func main()  {
	pic.Show(Pic)
}


//查看图片借助站长工具，base６４转换成图片