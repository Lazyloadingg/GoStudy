package main

//学习博客
// https://www.liwenzhou.com/posts/Go/12-interface/
import (
	"fmt"
)

type Cat struct{}

func (c Cat) Say() {
	fmt.Printf("\"喵喵喵\": %v\n", "喵喵喵")
}

func (c Cat) Eat() {
	fmt.Printf("\"吃东西\": %v\n", "吃东西")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Printf("\"汪汪汪\": %v\n", "汪汪汪")
}

type Sayer interface {
	Say()
	Eat()
}

// 接收实现了所有接口方法的类型
func AnimalSay(s Sayer) {
	fmt.Printf("\"Animal\": %v\n", "Animal")
	s.Say()
}

// 接口也可以组合起来生成新的接口
type Reader interface {
	read()
}
type Writer interface {
	write()
}
type ReadWriter interface {
	Reader
	Writer
}

type Person struct {
}

// 结构体中有其他结构体作为成员变量，可以实现类似`继承`的效果，调用成员结构体的方法函数
type Student struct {
	Person
}

func (p Person) read() {
	fmt.Printf("\"读书\": %v\n", "读书")
}
func (p Person) write() {

}

// 重写read，`Student`中的`Person`已经实现了`read`方法，那么如果`Student`再实现方法内部要写完整
func (s Student) read() {
	// s.read()不能这么写，会产生歧义导致报错
	s.Person.read() //如果内部要调用`Person`的`read`方法，必须写完整
}

// 空接口 空接口就是没有定义任何方法的接口，那么任何类型都可以认为是实现了空接口，因此空接口变量可以接受任意类型的值
type empty interface{}

// 作为函数参数，接收任意类型的值
func anyType(e empty) {
	fmt.Printf("任意类型e: %v %T\n", e, e)
}

func main() {
	//初看下来感觉go中的interface和iOS中的protocol好像，都是定义要实现方法，具体实现让遵循了接口和协议的类型去做

	//每种动物都要单独做同样的事情
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()

	//将Say抽象成接口，并用一个公共方法`AnimalSay`处理
	AnimalSay(c)
	//AnimalSay(d) //因为Sayer接口有两个参数，Dog类型只实现了一个，因此报错

	//接口类型也可以声明成变量，他可以接收所有实现了接口方法的类型
	var i Sayer
	i = c
	i.Say()

	var r Reader
	r = Person{}
	r.read()

	//组合类型接口变量只能接受实现了内部所有子接口方法的类型
	var rw ReadWriter
	rw = Person{}
	rw.write()

	//空接口 可以接受任意类型值，有点像iOS里的id类型，同样也可以作为函数参数，作为函数参数的话，函数就可以接收任意类型的参数
	var e empty
	e = 10 //接收整形
	fmt.Printf("e: %v\n", e)
	anyType(e)
	e = Person{} //接收结构体
	fmt.Printf("e: %v\n", e)
	anyType(e)
	e = []int{} //接收切片类型
	fmt.Printf("e: %v\n", e)
	anyType(e)

	//同样可以作为map的值类型
	m := make(map[string]interface{})
	m["name"] = "小明"             //名字string类型
	m["age"] = 18                //年龄int类型
	m["habit"] = []interface{}{} //爱好习惯切片类型，切片中又是任意类型
	fmt.Printf("m: %v\n", m)

	//空接口不能调用任何方法 否则会panic报错
	// var w ReadWriter
	// w.read()

	//接口类型变量因为可以接受任意实现了全部接口方法的类型，因此接口类型变量内部分为两部分，第一部分记录变量数据具体哪个类型，第二部分记录值

	// 结构体中有其他结构体作为成员变量，可以实现类似`继承`的效果，调用成员结构体的方法函数

	st := Student{}
	st.read()        //结构体`Student`中有匿名成员`Person`，`Person`实现了`read`函数，因此st可以直接通过`.`调用`read`函数
	st.Person.read() //上面是简略写法，这是完整写法

}
func test(a any) {

}
