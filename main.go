package main

import (
	"bytes"
	"fmt"
	"strings"
)

const PI = 3.1415926

func test() string {
	return "测试"
}

func main() {

	// // PI = 222
	// out = "222"
	// var name string = "小刚"
	// var age int = 19
	// s := user.Hello()
	// name2, _ := user.Info()
	// fmt.Printf("name2: %v\n", name2)
	// fmt.Printf("s: %v\n %d  %v", s, age, name)

	// fmt.Printf("out: %v\n", out)
	// fmt.Printf("PI: %v\n", PI)

	// fmt.Printf("%T\n", name)

	// a := 100
	// p := &a
	// fmt.Printf("p: %v\n", *p)
	// fmt.Printf("指针类型：%T\n", p)

	// // fmt.Printf("%T", test)

	// var fun = test
	// fmt.Printf("fun: %v\n", fun)

	// res := a >= 18
	// if res {
	// 	fmt.Printf("成年了")
	// } else {
	// 	fmt.Printf("成年了")
	// }

	// fmt.Printf("大小：%v\n", unsafe.Sizeof(a))
	// fmt.Printf("最大值：%v\n最小值:%v", math.MaxInt64, math.MinInt64)

	// test_string()
	// test_flow()

	// test_array()
	test_slice()
	// test_map()
	fmt.Printf("%v", test_repeatCount("how do you do"))
}

func test_string() {

	var s string = "hello 小明"
	var s1 = "hello 小航"
	s2 := "hello 小红"

	s3 := `	
	啦啦啦
	啦啦啦
	`
	fmt.Printf("%v\n%v\n%v\n%v\n", s, s1, s2, s3)
	//获取长度len
	fmt.Printf("长度：%v\n", len(s))

	//字符串拼接
	s4 := s + s1
	fmt.Printf("拼接后%v\n", s4)
	//格式化拼接
	s5 := fmt.Sprintf("%s和%v", s, s1)
	fmt.Printf("格式化拼接后%v\n", s5)

	//buffer写入
	var buffer bytes.Buffer
	buffer.WriteString("第一句")
	buffer.WriteString("第二句")
	fmt.Printf("buffer写入:%v\n", buffer.String())

	//切片，通过索引获取指定位置字符
	fmt.Printf("%v\n", s[0])   //104 单独位置获取的是原始的字符在ascii中的序号
	fmt.Printf("%c\n", s[0])   //h 这样表示就可以打印真正的字符
	fmt.Printf("%v\n", s[1:4]) //ell
	fmt.Printf("%v\n", s[2:])  //llo 小明
	fmt.Printf("%v\n", s[:6])  //hello

	//分割字符串
	fmt.Printf("分割字符串：%v\n", strings.Split(s, " "))

	//字符串比较
	fmt.Printf("字符串比较：%v\n", strings.EqualFold("233", "2332"))

	//其他常用字符串操作比如前后缀判断，是否包含指定字符等 可查看`strings`标准库
}

func test_flow() {

	a := [...]int{1, 3, 5, 7, 9}

	//区间循环
	for indx, v := range a {
		fmt.Println(indx, v)
	}

	test := 'A'

	switch {
	case test == 'A':
		println("是A啊")
	default:

	}

	b := [3]int{1, 2, 3}

	fmt.Printf("b: %v\n", b)

	for i := 0; i < 10; i++ {
		if i == 5 {
			//goto 无条件跳转
			goto LABEL
		}
	}

LABEL:
	print("跳转过来了")

}

func test_array() {
	var arr [2]int
	fmt.Printf("arr: %v\n", arr)

	var arr2 = [...]int{1, 2, 3}
	fmt.Printf("arr: %v\n", arr2)

	//访问数组元素编译期间会判断，如果越界则报错比如`arr2[3]`报错
	fmt.Printf("%v\n", arr2[2])
	print("长度：", len(arr2))

	count := len(arr2)
	//for循环遍历
	for i := 0; i < count; i++ {
		fmt.Printf("arr2[%v]: %v\n", i, arr2[i])
	}

	//for range遍历
	for index, v := range arr2 {
		fmt.Printf("arr2[index]: %v\n", arr2[index])
		fmt.Printf("v: %v\n", v)
	}

}
func test_slice() {
	//slice 切片，其实底层也是数组实现，不过数组长度不可变，切片可以自动扩容

	var s = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("s:%p %v\n", &s, s)
	s = append(s, 11)
	fmt.Printf("s:%p %v\n", &s, s)

	//声明切片方式1 ： 这种方式情况下切片分配内存未初始化，为空
	var s1 []int
	fmt.Printf("指针：%p , s1: %v\n", &s1, s1)

	//声明切片方式1 ： 这种方式情况下切片分配内存并初始化，为初始值
	var s2 = make([]int, 2)
	fmt.Printf("--指针：%p , s2: %v\n", &s2, s2)

	//通过数组方式初始化
	arr := [...]int{1, 2, 3, 4, 5, 6}
	s3 := arr[1:4]
	fmt.Printf("s3: %v\n", s3)

	//添加元素
	s3 = append(s3, 11)
	fmt.Printf("添加元素s3: %v  %p\n", s3, s3)

	//删除元素(系统没提供直接删除方法，好傻吊)，这种方式相当于将除了要删除的元素之外的其他元素加入数组
	s3 = append(s3[:2], s3[3:]...)
	fmt.Printf("删除元素s3: %v  %p\n", s3, s3)

	//这种形式类似浅拷贝
	var s4 = s3
	fmt.Printf("s4: %p  %v\n", s4, cap(s4))
	fmt.Printf("s3: %p  %v\n", s3, cap(s4))
	//修改s4的元素，s3也会被修改
	s4[0] = 100
	s4 = append(s4, 99)
	fmt.Printf("s4: %p  %v  %v\n", s4, cap(s4), s4)
	fmt.Printf("s3: %p  %v  %v\n", s3, cap(s3), s3)

	//深拷贝
	s5 := make([]int, len(s4))
	copy(s5, s4)
	fmt.Printf("s5: %v\n", s5)

}

func test_map() {
	//通过make函数初始化
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	m1["name"] = "小明"
	fmt.Printf("m1: %v\n", m1)

	//直接初始化
	m2 := map[string]string{
		"name": "小刚",
		"age":  "18",
	}
	fmt.Printf("m2: %v\n", m2)

	//判断是否存在某个Key，通过对某个key取值，取到就是存在否则不存在
	v, ok := m2["name"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

	v1, ok1 := m2["height"]
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("ok1: %v\n", ok1)

	//删除键值对
	delete(m2, "age")
	fmt.Printf("删除后m2: %v\n", m2)
	//遍历map，forrange方式，针对map就是key,value，针对数组就是index,value
	for k, v := range m2 {
		fmt.Printf("遍历k: %v\n", k)
		fmt.Printf("遍历v: %v\n", v)
	}

	for dict := range m2 {
		fmt.Printf("dict: %v\n", dict)
	}

}

func test_repeatCount(text string) (res map[string]int) {

	arr := strings.Split(text, " ")
	res = make(map[string]int)
	for _, v := range arr {
		_, ok := res[v]
		if ok {
			res[v] = res[v] + 1
		} else {
			res[v] = 1
		}
	}

	return res
}
