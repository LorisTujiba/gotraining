package main

import "fmt"

func main(){

	/*==============
		Single Loop
		============
	*/
	for i:=0;i<10;i++{
		fmt.Print(i)
	}

	fmt.Println()

	/*==============
		Nested Loop
		============
	*/
	for i:=0;i<5;i++{
		for j:=0;j<10;j++  {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*======================
		Loop with condition
		====================
		No "while" in GO
		so this is the
		alternative
	*/
	i:=0
	for i < 10{
		fmt.Print(i)
		i++
	}

	fmt.Println()

	j:=0
	for {
		fmt.Print(j)
		if j >= 9{
			break
		}
		j++
	}

	fmt.Println()

	/*=====================
		Break and Continue
		===================
	*/
	k:=0

	for {
		k++
		if k % 2 ==1{
			continue
		}
		fmt.Print(k)

		if k > 10 {
			break
		}
	}
}
