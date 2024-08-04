package main

import (
	f "fmt"
	"strings"
)

func main() {
	contains()
	prefix()
	suffix()
	count()
	index()
	replace()
	repeat()
	split()
	join()
	lowerUpper()
}

func contains() {
	var isExist = strings.Contains("john wick", "john")
	f.Println(isExist)
}

// mendeteksi string yang diawali string tertentu
func prefix() {
	var isPrefix1 = strings.HasPrefix("john wick dodo", "jo")
	f.Println(isPrefix1)

	var isPrefix2 = strings.HasPrefix("john wick dodo", "do")
	f.Println(isPrefix2)
}

// mendeteksi string yang diakhiri string tertentu
func suffix() {
	var isSuffix1 = strings.HasSuffix("anime", "im")
	f.Println(isSuffix1)

	var isSuffix2 = strings.HasSuffix("anime", "me")
	f.Println(isSuffix2)
}

// menghitung jumlah string tertentu
func count() {
	var howMany = strings.Count("ethan hunt", "t")
	f.Println(howMany)
}

// mencari posisi index
func index() {
	var index1 = strings.Index("ethan hunt", "ha")
	f.Println(index1)

	var index2 = strings.Index("ethan hunt", "n")
	f.Println(index2)
}

func replace() {
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	f.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	f.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	f.Println(newText3)
}

func repeat() {
	var str = strings.Repeat("na", 4)
	f.Println(str)
}

// untuk memisahkan string
func split() {
	var string1 = strings.Split("the dark knight", " ")
	f.Println(string1)

	var string2 = strings.Split("batman", "")
	f.Println(string2)
}

func join() {
	var data = []string{"banana", "papaya", "tomato"}
	var str = strings.Join(data, "-")

	f.Println(str)
}

func lowerUpper() {
	var str1 = strings.ToLower("AMBasiNG")
	var str2 = strings.ToUpper("ambaSIng")
	f.Println(str1)
	f.Println(str2)
}
