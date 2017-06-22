package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

/*
https://www.hackerrank.com/challenges/ctci-ransom-note

Input Format

The first line contains two space-separated integers describing the respective values of  (the number of words in the magazine) and  (the number of words in the ransom note).
The second line contains  space-separated strings denoting the words present in the magazine.
The third line contains  space-separated strings denoting the words present in the ransom note.

Constraints

Each word consists of English alphabetic letters (i.e.,  to  and  to ).
The words in the note and magazine are case-sensitive.
Output Format

Print Yes if he can use the magazine to create an untraceable replica of his ransom note; otherwise, print No.*/


var magazineLen, ransomLen, used int
var magazine, ransom string
var splitMagazine, splitRansom []string
var magazineBuckets = make(map[int]string)

func main(){


	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader,&magazineLen)
	fmt.Fscan(reader,&ransomLen)
	reader.ReadString('\n')
	magazine,_ = reader.ReadString('\n')
	ransom,_ = reader.ReadString('\n')

	splitMagazine = strings.Fields(magazine)
	splitRansom = strings.Fields(ransom)

	//Index words based on the first letter
	for i:= 0; i<magazineLen;i++{
		magazineBuckets[hash(splitMagazine[i])] = splitMagazine[i]
	}

	check()

}

func hash(word string) int{
	return int(word[0])-65
}
func check(){
	for i:=0; i<ransomLen; i++{
		value := hash(splitRansom[i])
		if len(magazineBuckets[value])!=0 {
			for j := 0; j < len(magazineBuckets[value]); j++ { // loop based on the corresponding initial
				if splitRansom[i] == magazineBuckets[value][j] {
					fmt.Println("asda",len(magazineBuckets[value]))
					used++
					magazineBuckets[value][j] = ""
					fmt.Println("asda",len(magazineBuckets[value]))
					break
				}
			}
		}else{
			break
		}
	}
	if used == ransomLen{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}

}