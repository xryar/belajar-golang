package main

import (
	f "fmt"
	"regexp"
)

var text = "banana burger soup"
var regex, _ = regexp.Compile(`[a-z]+`)

func main() {
	applyRegex()
	matchString()
	findString()
	findAllString()
	replaceAllString()
	replaceAllStringFunc()
	regexSplit()
}

func applyRegex() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		f.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	f.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	f.Printf("%#v \n", res2)
}

func matchString() {
	var isMatch = regex.MatchString(text)
	f.Println(isMatch)
}

func findString() {
	var idx = regex.FindStringIndex(text)
	f.Println(idx)

	var str = text[0:6]
	f.Println(str)
}

func findAllString() {
	var str1 = regex.FindAllString(text, -1)
	f.Println(str1)

	var str2 = regex.FindAllString(text, 1)
	f.Println(str2)
}

func replaceAllString() {
	var str = regex.ReplaceAllString(text, "potato")
	f.Println(str)
}

func replaceAllStringFunc() {
	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})

	f.Println(str)
}

func regexSplit() {
	var regex, _ = regexp.Compile(`[a-b]+`)

	var str = regex.Split(text, -1)
	f.Printf("%#v \n", str)
}
