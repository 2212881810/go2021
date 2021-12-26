package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名字叫:" + name)
}

func main() {
	u := User{
		Name: "郑钦锋",
		Age:  100,
	}

	s := Student{
		Class: "一年级",
		User:  u,
	}

	//check(u)
	//testReflect(u)
	//testReflect2(s)

	fmt.Println(s)
	// 参数:指针类型
	testReflect3(&s)
	fmt.Println(s)

	testReflect4(u)

}

/**
通过反射调用对象的方法
*/
func testReflect4(inter interface{}) {
	v := reflect.ValueOf(inter)

	method := v.MethodByName("SayName")
	//method := v.Method(0)  // 通过index调用方法

	// reflect.ValueOf()会创建一个Value实例
	method.Call([]reflect.Value{reflect.ValueOf("admin")})

}

/**
通过反射修改原始数据的值，注意入参一定要是指针类型
*/
func testReflect3(inter interface{}) {
	v := reflect.ValueOf(inter)
	e := v.Elem()
	e.FieldByName("Class").SetString("二年级")
	fmt.Println(inter)

	t := reflect.TypeOf(inter)
	// 通过kind方法获取inter的类型
	tk := t.Kind()
	if tk == reflect.Struct {
		fmt.Println("我是struct")
	}

	if tk == reflect.String {
		fmt.Println("我是String")
	}
}

/**
通过反射，按层级取属性值
*/
func testReflect2(inter interface{}) {
	// 获取类型
	t := reflect.TypeOf(inter)
	// 获取值
	v := reflect.ValueOf(inter)
	fmt.Println(t, v)
	fmt.Println(v.FieldByName("Class"))

	// 按index取值， 可以按层级取值
	fmt.Println(v.FieldByIndex([]int{0}))
	fmt.Println(v.FieldByIndex([]int{1, 0}))
	fmt.Println(v.FieldByIndex([]int{1, 1}))
}

/**
遍历属性值
*/
func testReflect(inter interface{}) {
	// 获取类型
	t := reflect.TypeOf(inter)
	// 获取值
	v := reflect.ValueOf(inter)

	// 有多少个属性
	fields := t.NumField()

	for i := 0; i < fields; i++ {
		// 遍历，获取每个属性的值
		fmt.Println(v.Field(i))
	}

	//fmt.Println(t, v)

}

/**
断言
*/
func check(inter interface{}) {
	switch inter.(type) {
	case User:
		fmt.Println("我是User")
		break
	case Student:
		fmt.Println("我是Student.")
		break
	}
}
