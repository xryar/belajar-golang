package main

import (
	f "fmt"
	"strconv"
)

func main() {
	stringToInt()
	intToString()
	parseInt()
	formatInt()
	stringToFloat()
	floatToString()
	stringToBoolean()
	booleanToString()
	withCasting()
	castingStringToByte()
	typeAssertions()
}

func stringToInt() {
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		f.Println(num)
	}
}

func intToString() {
	var num = 124
	var str = strconv.Itoa(num)

	f.Println(str)
}

func parseInt() {
	var str = "1010"
	var num, err = strconv.ParseInt(str, 2, 8)

	if err == nil {
		f.Println(num)
	}
}

func formatInt() {
	var num = int64(24)
	var str = strconv.FormatInt(num, 8)

	f.Println(str)
}

func stringToFloat() {
	var str = "24.12"
	var num, err = strconv.ParseFloat(str, 32)

	if err == nil {
		f.Println(num)
	}
}

func floatToString() {
	var num = float64(24.12)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	f.Println(str)
}

func stringToBoolean() {
	var str = "true"
	var bul, err = strconv.ParseBool(str)

	if err == nil {
		f.Println(bul)
	}
}

func booleanToString() {
	var bul = true
	var str = strconv.FormatBool(bul)

	f.Println(str)
}

func withCasting() {
	var a float64 = float64(24)
	f.Println(a)

	var b int32 = int32(24.00)
	f.Println(b)
}

func castingStringToByte() {
	var text1 = "halo"
	var b = []byte(text1)

	f.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	f.Println(" ")

	//byte to string
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	f.Printf("%s \n", s)
}

func typeAssertions() {
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	for _, val := range data {
		switch val.(type) {
		case string:
			f.Println(data["nama"].(string))
		case int:
			f.Println(data["grade"].(int))
		case float64:
			f.Println(data["height"].(float64))
		case bool:
			f.Println(data["isMale"].(bool))
		case []string:
			f.Println(data["hobbies"].([]string))
		default:
			f.Println(val.(int))
		}
	}

}
