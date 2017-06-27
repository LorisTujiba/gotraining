package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct{
	lat,long string
	err error
}

func (n *NorgateMathError) Error() string{
	return fmt.Sprintf("Error occured : %v %v %v",n.lat,n.long,n.err)
}

func main(){
	_,err := Square(-10)
	if err!=nil{
		log.Println(err)
	}
}

func Square(f float64) (float64,error){
	if f < 0{
		nme := fmt.Errorf("Error : %v",f)
		return 0, &NorgateMathError{"Ayam","Kambing",nme}
	}
	return 42,nil
}