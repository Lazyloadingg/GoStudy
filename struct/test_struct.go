package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//定义两类型，一种全新类型和swift类似
type newInt int
type fn func(int, int) int

//起别名，本质还是后边的类型，只是换了个名字
type myInt = int

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	name string
}

//go中没有类，同样也没有继承，但是可以通过给结构体添加匿名变量的方式，达到"继承"的效果
type Dog struct {
	age    int
	Animal //匿名变量，只有类型没有名字，将来Dog的实例也可以调用Animal的方法
}

/*
结构体内字段想要在其他包中访问就要将首字母大写，如果小写就只能在当前包访问
这时如果json原始数据字段名为小写，或者字段名完全不同，就要通过指定tag，实现json序列化该字段时的key
*/
type Student struct {
	//字段首字母大写，但是我们指定json中小写字母的key也可以解析
	Name string `json:"name"`
	Age  int
}
type Class struct {
	Title    string `json:"title"`
	Students []*Student
}

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func main() {

	var a newInt = 10
	var b myInt = 5
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	//使用键值对初始化
	p1 := Person{
		Name: "小明",
		Age:  19,
	}

	fmt.Printf("person: %v\n", p1)

	var p2 Person
	fmt.Printf("p2: %v\n", p2)

	p3 := &Person{
		Name: "小刚",
		Age:  20,
	}
	fmt.Printf("p3: %v\n", p3)
	p4 := &p1
	fmt.Printf("p4: %v\n", p4)
	p4.Name = "2333"
	fmt.Printf("--p4: %v\n", p4)
	fmt.Printf("p1: %v\n", p1)
	p1.Name = "666"
	fmt.Printf("--p4: %v\n", p4)
	fmt.Printf("p1: %v\n", p1)

	p5 := initialized_struct("小红", 17)
	fmt.Printf("p5: %v\n", p5)
	p5.smile()
	s := strings.Split("233", "")
	fmt.Printf("s: %v\n", s)

	var n newInt = 10
	fmt.Printf("n.sum(3): %v\n", n.sum(3))

	p6 := Animal{
		name: "狗",
	}
	fmt.Printf("p6: %v\n", p6)
	p6.run()

	p7 := Dog{
		age: 3,
	}
	p7.eat()
	p7.name = "旺财"
	p7.run() //run方法是Animal结构体的，此时Dog也可调用

	//初始化一个班级
	class := Class{
		Title:    "向日葵班",
		Students: []*Student{},
	}
	fmt.Printf("class: %v\n", class)

	for i := 0; i < 10; i++ {
		class.Students = append(class.Students, &Student{Name: strconv.Itoa(i+1) + "号", Age: 18})
	}

	fmt.Printf("class: %v\n", class)

	//结构体转json
	data, err := json.Marshal(class)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("data: %v\n", string(data))

	//json转结构体

	jsonStr := "{\"title\":\"向日葵班1\",\"Students\":[{\"Name\":\"1号\",\"Age\":18}]}"
	c1 := Class{}

	bytes := []byte(jsonStr)
	fmt.Printf("bytes: %v\n", bytes)
	err1 := json.Unmarshal(bytes, &c1)
	if err1 != nil {
		fmt.Printf("\"错了\": %v\n", err1)
	}
	fmt.Printf("json to struct c1: %v\n", c1)

}

// go结构体没有swift强大，没有构造器，因此自己手动写一个函数实现，因为go结构体是值类型，传递是copy，频繁操作消耗性能,因此返回结构体指针
func initialized_struct(name string, age int) *Person {

	res := Person{
		Name: name,
		Age:  age,
	}
	return &res
}

//go没有像c++ swift那样的类，如果让某些方法值针对特定类型变量，可以定义方法，方法整体和函数类似，只不过在函数名前面加了个接收者类型
func (p Person) smile() {
	fmt.Printf("p: %v 哈哈哈\n", p)
}

//可以给任意类型添加方法，举例给我们自定义的类型newInt添加sum方法
func (m newInt) sum(a int) int {
	return a + 1 + int(m)
}

func (a Animal) run() {
	fmt.Printf("a.name: %v跑了\n", a.name)
}
func (d Dog) eat() {
	fmt.Printf("d: %v岁的狗在吃饭\n", d.age)
}
