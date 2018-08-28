package main

import (
	"fmt"
)

func main() {

	/**
	切片有容量(capacity)和长度(length)两个属性
	*/

	/*
	 用make函数定义切片x，这个时候x的容量是5，长度也是5。
	*/
	var x = make([]float64, 5)
	fmt.Println("Capacity:", cap(x), "Length:", len(x))

	/**
	使用make函数定义了切片y，这个时候y的容量是10，长度是5。
	虽然切片的容量可以大于长度，但是赋值的时候要注意最大的索引仍然是len(x)－1。否则会报索引超出边界错误。
	*/
	var y = make([]float64, 5, 10)
	fmt.Println("Capacity:", cap(y), "Length:", len(y))

	for i := 0; i < len(y); i++ {
		y[i] = float64(i * 2)
	}

	fmt.Printf("y: ", y)

}
