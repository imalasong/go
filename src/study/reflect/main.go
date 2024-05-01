package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//testA()

	testS()
}

func testS() {
	s := Stud{
		Id:   10,
		Name: "imalasong",
		desc: "hello",
	}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(&s).Elem()

	fmt.Printf("kind:%v,name:%v,value:%v,fileCount:%d,methodCount:%d\n", t.Kind(), t.Name(), v, v.NumField(), v.NumMethod())

	for i := 0; i < t.NumField(); i++ {
		fieldV := v.Field(i)
		fieldT := t.Field(i)
		if fieldV.CanSet() {
			switch fieldV.Kind() {
			case reflect.String:
				fieldV.Set(reflect.ValueOf(fieldV.Interface().(string) + ",hello"))
			case reflect.Int64:
				fieldV.Set(reflect.ValueOf(fieldV.Interface().(int64) + 1))
			}
		}

		//modify unexport filed
		if fieldT.Name == "desc" {
			*(*string)(unsafe.Pointer(v.UnsafeAddr() + fieldT.Offset)) = "helloaaaaa"
		}

		//get tag
		tagC := fieldT.Tag.Get("c")

		fmt.Printf("name=%v,type=%v,value=%v,tagC=%v\n", fieldT.Name, fieldV.Type(), fieldV, tagC)
	}

	fmt.Printf("new Value:%v\n", s)

	fmt.Println("--------------------------------------------")

	//method

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		methodV := v.Method(i)
		fmt.Printf("call methodName:%v,export:%v\n", method.Name, method.IsExported())
		args := make([]reflect.Value, 0)
		methodV.Call(args)
	}
}

func testA() {
	var f float64 = 10

	t := reflect.TypeOf(f)
	t2 := reflect.TypeOf(&f)
	v := reflect.ValueOf(f)
	v2 := reflect.ValueOf(&f)

	f2 := v.Interface().(float64)
	f21 := v2.Interface().(*float64)

	//修改值
	elem := v2.Elem()
	fmt.Println(elem.CanSet())
	elem.Set(reflect.ValueOf(float64(19)))

	fmt.Println(f)
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(v)
	fmt.Println(v2)
	fmt.Println(f2)
	fmt.Println(f21)
}

type Stud struct {
	Id   int64  `c:"id"`
	Name string `c:"name"`

	desc string
}

func (s Stud) A1() {
	fmt.Println("hhllo A1")
}
func (s *Stud) A2() {
	fmt.Println("hhllo A2")
}

func (s Stud) a1() {
	fmt.Println("hhllo a1")
}
