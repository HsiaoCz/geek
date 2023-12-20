package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	var c *float32
	var d myInt // 自定义类型
	var e rune  // 类型别名
	reflectType(c)
	reflectType(d)
	reflectType(e)

	type person struct {
		name string
		age  int
	}

	type book struct{ title string }

	var f = person{
		name: "沙河小王子",
		age:  18,
	}
	var g = book{title: "《hello》"}
	reflectType(f)
	reflectType(g)

	var s float32 = 3.14
	reflectValue(s)

	var bb int = 100
	reflectValue(bb)

}

func reflectType(x any) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x any) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Println(int64(v.Int()))
	case reflect.Float32:
		fmt.Println(float32(v.Float()))
	}
}

func ReflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func ReflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
