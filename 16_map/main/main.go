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

func main(){

	//male(map[type of the key]type of the value)
	// "{}" is called composite literal
	//var greeting map[string]string //we get nill reference
	//var greeting = map[string]string{}#1
	var greeting = make(map[string]string)//#2

	greeting["Edward"] = "Good Morning"// key is edward
	greeting["Udin"] = "Selamat Pagi"//key is udin and the value is selamat pagi

	fmt.Println(greeting)

	anotherGreeting := map[string]string{"John":"Soldier","Harley":"Doctor"}

	fmt.Println(anotherGreeting["John"])//call map where having john as a key

	//integer key map
	var instruments = map[int]string{
		1:"guitar",
		2:"bass",
		3:"drum",
		4:"keyboard",
		5:"Saxophone",
	}
	fmt.Println(instruments)
	delete(instruments,5)//delete the value that has 5 as the key

	//delete(instruments,2)

	if val,exist := instruments[2];exist{ //exist is comma ok idiom
		fmt.Println(val)
		fmt.Println(exist)
	}else {
		fmt.Println("it doesnt exist")
	}

	//map inside a map
	var mapInside = map[int] map[string]string{
		1:map[string]string{
			"name":"Josh",
			"origin":"US",
		},
		2:map[string]string{
			"name":"Claire",
			"origin":"Canada",
		},
	}

	fmt.Println(mapInside)

	//map traversing, can be done using for range

	for index,instrument := range instruments{
		fmt.Println("Index : ",index,"Instrument : ",instrument)
	}

}