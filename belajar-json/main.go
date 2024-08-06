package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
}

var jsonString = `{"Name": "john wick", "Age": 27}`
var jsonData = []byte(jsonString)

func main() {
	decodeJsonToStruct()
	decodeJsonToInterface()
	decodeArrayJsonToObject()
	encodeObjectToJsonString()
}

func decodeJsonToStruct() {
	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.Fullname)
	fmt.Println("age  :", data.Age)
}

func decodeJsonToInterface() {
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodeData = data2.(map[string]interface{})
	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age  :", decodeData["Age"])
}

func decodeArrayJsonToObject() {
	var jsonString = `[
		{"Name": "John wick", "Age": 27},
		{"Name": "Arya Rizki", "Age": 21}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User 1:", data[0].Fullname)
	fmt.Println("User 2:", data[1].Fullname)
}

func encodeObjectToJsonString() {
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
