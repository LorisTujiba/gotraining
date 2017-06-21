package main

import "fmt"

/*==============================
Maps
============================
Key Value Storage. If not
initialized, the value
is nil. unordered. it
is referenced type.
*/

func main() {

	//male(map[type of the key]type of the value)

	//var greeting map[string]string //we get nill reference
	//var greeting = map[string]string{}#1
	var greeting = make(map[string]string) //#2

	greeting["Edward"] = "Good Morning" // key is edward
	greeting["Udin"] = "Selamat Pagi"   //key is udin and the value is selamat pagi

	fmt.Println(greeting)

	anotherGreeting := map[string]string{"John": "Soldier", "Harley": "Doctor"}

	fmt.Println(anotherGreeting["John"]) //call map where having john as a key

	//integer key map
	var instruments = map[int]string{
		1: "guitar",
		2: "bass",
	}
	fmt.Println(instruments)
}
