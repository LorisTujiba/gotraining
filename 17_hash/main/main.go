package main

import (
	"bufio"
	"fmt"
	"github.com/LorisTujiba/gotraining/03_scope/values"
	"io/ioutil"
	"log"
	"net/http"
)

/*================================
Hash Table
================================
Data structure used to implement
an associative array. structure
that has key and values. these
values will be stored into
things called buckets.
where the desired
value can be
found.
*/

func main() {
	//Hashing Example
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt") //get Resource
	if err != nil {
		log.Fatal(err)
	}

	//scan pages
	scanner := bufio.NewScanner(res.Body)
	//defer res.Body.Close()
	scanner.Split(bufio.ScanWords) //splitting by words

	buckets := make([]int, 200) //create slice to hold counts

	for scanner.Scan() { //loop the words
		n := hashBucket(scanner.Text())
		//n := hashBucketDividebyTwelve(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123]) /*print from capital A to z
	return 1790 1260 910
	it means, there's 1790 words start with capital A
	1260 start with capital B and so on*/

	/*=============
	Read All
	=============
	*/

	res2, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res2.Body) //read all moby dick
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)

	/*=====================
	Building a Hash table
	=====================
	*/

	res3, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt") //get Resource
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(res3.Body)

	//defer res.Body.Close()
	scanner.Split(bufio.ScanWords) //splitting by words

	buckets2 := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets2 = append(buckets2, []string{})
	}

	for scanner.Scan() {
		word := scanner.Text() //getting each word
		n := hashBucketDividebyTwelve(word)
		buckets2[n] = append(buckets2[n], word)
	}

	//getting the amount of words on each bucket
	for i := 0; i < 12; i++ { //loop all of the 12 buckets
		fmt.Println(i, " - ", len(buckets2[i]))
		values.PrintSeparation("=", 60)
	}

	//see inside the bucket
	fmt.Println(buckets2[10])

}

func hashBucket(word string) int {
	asciiValue := int(word[0]) //get the first letter and convert to number(Ascii)
	return asciiValue
}

func hashBucketDividebyTwelve(word string) int { //break down into 12 buckets
	letter := int(word[0])
	buckets := letter % 12
	return buckets
}
