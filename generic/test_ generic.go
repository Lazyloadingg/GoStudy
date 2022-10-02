package main

import "fmt"

type TestInt interface {
	int8 | int16 | int32 | int64
}

func main() {
	// 1.18开始可以用`any`替换`interface{}`的写法，这样在用泛型是更简洁
	test1("interface写法")
	test2("any写法")
	res := test3[int8](10)
	fmt.Printf("res: %v\n", res)
	fmt.Printf("test4[int16](10, 8): %v\n", test4[int16](10, 8))
}

// interface写法
func test1(a interface{}) {
	fmt.Printf("a: %v\n", a)
}

// any写法，其实any是interface的别名，点进去看会发现`type any = interface{}`定义
func test2(a any) {
	fmt.Printf("a: %v\n", a)
}

// 泛型写法
func test3[T int8 | int16](a T) T {
	return a + 10
}

// test3写法没问题，但是一旦用于约束类型的形参列表很长的话，写起来就很麻烦，因此可以抽出来一个泛型约束类型`TestInt`
func test4[T TestInt](a, b T) T {
	return a + b
}
