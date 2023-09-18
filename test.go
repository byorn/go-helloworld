package main

import "fmt"

func tt() {
	fmt.Println("new file")
}

type mytest string

func (m mytest) print() {
	fmt.Println(m)
}

type friend []string

/*
* (f friend) is a receiver on a function
attach print to the type friend

f is the instance

by convention we are supposed to use only one letter for the instace
f is the self,or the this keyword..
*
*/
func (f friend) print() {

	for i, friendName := range f {
		fmt.Println(i, friendName)
	}
	fmt.Println("finished iterating friend names")
}

func testMaps1() {
	myMap := make(map[string]string)

	myMap["name"] = "byorn"

	fmt.Println(myMap)

	fmt.Println(myMap["name"])

	delete(myMap, "name")

	fmt.Printf("now name is %v\n", myMap["name"])
}

func testMaps2() {
	myMap := map[string]string{
		"name": "james",
		"age":  "12",
	}

	myMap["address"] = "asfd"

	fmt.Println(myMap)

	printMap(myMap)
}

func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("ok", key, value)
	}
}

/** interfaces **/
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

type srilankanBot struct{}

func (englishBot) getGreeting() string {
	return "hi there"
}

func (spanishBot) getGreeting() string {
	return "hola"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(sp spanishBot) {
//	fmt.Println("hola")
//}
