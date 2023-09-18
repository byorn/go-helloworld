package main

import (
	"fmt"
	"github.com/byorn/go-helloworld/util"
	util2 "github.com/byorn/test-go-modules/util"
)

func main() {

	//cards := newDeck()

	//hand, remainingDeck := deal(cards, 3)

	//hand.print()
	//remainingDeck.print()

	//saveToFile(cards, "firstfile.txt")
	//tt()
	//var m mytest = "helloworld"
	//m.print()
	//cards := createNewDeckFromFile("firstfile.txt")
	//cards.shuffle()
	//cards.print()
	//fmt.Println("----------")
	//cards.shuffle()
	//cards.print()

	//var p = newPerson("john", "adams")

	/** pointer = p  means value is copied, copy of the address location value **/
	/** pointer = &p means get original address location of p and pass it pointer **/
	//var pointer = &p
	//
	//p.updateLastNameWithOriginalAddressLocationPointer("bulocks")
	//fmt.Println(p)
	//fmt.Println(pointer)

	//testMaps1()
	//testMaps2()
	//eb := englishBot{}
	//printGreeting(eb)
	//sb := spanishBot{}
	//printGreeting(sb)

	//testFunc()

	//testGoRoutinesAndChannels()
	fmt.Println(util.DoSomethingFunkyUtilMethodInAnotherPackage())
	fmt.Println(util2.ThisIsSomeMethodToBeCalledFromAnotherRepository())
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
	fmt.Println(namePointer)
}
