package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*=============================================================
JSON | Marshal - Unmarshal & Encode - Decode
=============================================================
JSON is an open standard format that uses human-readable text
to transmit data objects consisting of attribute and value.
To create or unpack JSON, we can use Marshal/Encode and
Unmarshal/decode. Both do the same, except for:
	-Marshal does it to a string
	-Encode does it to a stream,
	such as getting a value
	from an API
*/

/*=================
Marshal - Example
=================
*/
type Hero struct {
	Name string //must be exported, if its not then it will not shows up
	Hp   int    `json:"-"`                      //using tags, telling json to exclude this property
	Mp   int    `json:"change the key to this"` //tags also
}

func main() {

	names := []string{"John", "Adam", "Eve"}

	fmt.Println(names)
	bs, _ := json.Marshal(names)

	fmt.Println(bs) // 91 34 74 111 ....
	//91 is Ascii for J, so we can see that its translating
	fmt.Println(string(bs)) //conversion to string, get the actual json

	var jakiro = Hero{
		"Jakiro",
		550,
		400,
	}
	byteSlice, _ := json.Marshal(jakiro)
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))

	/*===================
	  Unmarshal - Example
	  ===================
	*/

	supposedThisIsJonFromAnAPI := []byte(`["Mogul Khan","Lanaya","Fransiskarosa"]`)
	supposedThisIsJonFromAnAPI2 := []byte(`{"Name":"Mogul Khan","Hp":800,"Mp":250}`)

	json.Unmarshal(supposedThisIsJonFromAnAPI, &names)   //unmarshall to names addres
	json.Unmarshal(supposedThisIsJonFromAnAPI2, &jakiro) //unmarshall to names addres

	fmt.Println(names)
	fmt.Println(jakiro)

	/*===================
	  Encoding - Example
	  ===================
	*/
	json.NewEncoder(os.Stdout).Encode(jakiro)

	/*===================
	  Decoding - Example
	  ===================
	*/

	var reader = strings.NewReader(`{"Name":"Ray Toro","Hp":500,"Mp":400}`)
	json.NewDecoder(reader).Decode(&jakiro)
	fmt.Println(jakiro.Name)
	fmt.Println(jakiro.Hp)
	fmt.Println(jakiro.Mp)

}
