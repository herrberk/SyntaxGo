package main

import ("fmt"
	//"math"
	"strings"
	"sort"
	//"os"
	//"log"
	//"io/ioutil"
	//"strconv"
)
func main() {
	sample := "Hello World"

	fmt.Println(strings.Contains(sample,"orl"))
	fmt.Println(strings.Count(sample,"l"))
	fmt.Println(strings.Index(sample,"llo"))
	fmt.Println(strings.Replace(sample,"l","x",2))

	csvString := "s,a,c,t,k,d"
	fmt.Println(strings.Split(csvString,","))
	listOfLetters := [] string{"c","a","b"}

	sort.Strings(listOfLetters)

	fmt.Println("Letters: ",listOfLetters)

	listOfNums := strings.Join([]string{"3","2","1","0"}, ",")

	fmt.Println(listOfNums)

}
