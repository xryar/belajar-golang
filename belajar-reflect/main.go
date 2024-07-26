package main

import (
	f "fmt"
	r "reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = r.ValueOf(s)

	if reflectValue.Kind() == r.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f.Println("nama      :", reflectType.Field(i).Name)
		f.Println("tipe data :", reflectType.Field(i).Type)
		f.Println("nilai     :", reflectValue.Field(i).Interface())
		f.Println(" ")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = r.ValueOf(number)
	var s1 = &student{Name: "wick", Grade: 2}
	var s2 = &student{Name: "john wick", Grade: 2}

	f.Println("nama :", s2.Name)

	var reflectValue2 = r.ValueOf(s2)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]r.Value{
		r.ValueOf("wick"),
	})
	f.Println("nama :", s2.Name)
	f.Println(" ")

	s1.getPropertyInfo()
	f.Println("tipe  variabel :", reflectValue.Type())
	f.Println("nilai variabel :", reflectValue.Interface().(int))
}
