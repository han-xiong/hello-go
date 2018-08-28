package main

import "fmt"

func main() {
	// 定义一个切片， 容量为10， 长度为5
	var arr = make([]int, 5, 10)
	fmt.Println("cap: ", cap(arr), "len: ", len(arr))

	// 当切片扩充
	arr = append(arr, 5, 6, 7, 8, 9, 10)
	fmt.Println("Capacity:", cap(arr), "Length:", len(arr))
	fmt.Println(arr)

	// 修改值
	arr[1] = 10
	fmt.Println(arr)
}
