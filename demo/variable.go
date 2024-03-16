package main

import "fmt"

/**
go 里面变量声明了一定要使用
 */
func main() {
	// 先声明 再赋值
	var name string
	name = "niyulu";
	fmt.Println(name)
	// 声明并赋值
	var age int = 18
	fmt.Println(age)
	// 省略类型
	var sex = "男"
	fmt.Println(sex)
	// 直接推断
	tag := "帅"
	fmt.Println(tag)
}
