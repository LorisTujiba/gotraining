package main

import (
"log"
"errors"
"fmt"
)

func main(){
	_, err := square(-10)
	if err != nil{
		log.Fatal(err)
	}
}

func square(f float64) (float64,error){//multi return values
	if f < 0{
		ErrNorgateMath := fmt.Errorf("Error : %v",f)
		return 0, ErrNorgateMath
	}
	return 42,nil
}
