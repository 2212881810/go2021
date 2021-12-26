package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var current reflect.Value

	var st = Student{
		Name: "admin",
		Age:  18,
	}

	//fs := reflect.ValueOf(st).NumField()

	t := reflect.TypeOf(st)

	// 也能获取所有的属性值
	//fsNums := t.NumField()

	fmt.Printf("TypeOf is :%v\n", t)
	current = reflect.ValueOf(st)

	fmt.Printf("valueof's result is: %v\n", current)

	typ := current.Type()
	fmt.Printf("current type is: %v\n", typ)

	numFields := current.NumField()
	fmt.Println(numFields)

	var fld reflect.StructField
	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)
		fmt.Printf("%d field is: %v\n", i, fld)

		tagJsonValue := fld.Tag.Get("json")

		fmt.Printf("Get Tag value : %v\n", tagJsonValue)

	}

}
