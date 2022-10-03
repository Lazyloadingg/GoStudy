package main

import (
	"fmt"
)

// 1.18之前接口仅支持定义方法集，1.18之后，为了更好地支持泛型，接口增加支持了类型集，语法如下
type TestInt interface {
	//类型前加波浪线`~`表示支持通过此类型派生出来的任何类型，比如`type MyInt int` 其中的`MyInt`就是`int`的派生类型，
	//这样的类型可以有很多，我们不能每增加一个就改一次泛型约束条件
	~int | int8 | int16 | int32 | int64
}

// 除了通过interface约束泛型类型集合，也可以自定义类型约束泛型集合
type MyInt[T int | int64 | string] int

func main() {
	// 1.18开始可以用`any`替换`interface{}`的写法，这样在用泛型是更简洁
	test1("interface写法")
	test2("any写法")
	res := test3[int8](10)
	fmt.Printf("res: %v\n", res)
	fmt.Printf("test4[int16](10, 8): %v\n", test4(10, 8))
	fmt.Printf("test5[int](13, 12): %v\n", test5(13, 12))

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

// 除了通过接口约束泛型类型集合，还可以自定义类型，比如`MyInt`这种写法
func test5[T MyInt[int]](a, b T) T {
	return a + b
}
