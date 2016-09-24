package main

import "fmt"

func main() {

	//Primitive Types and Variable Declarations

	fmt.Printf("Hello,\n")
	fmt.Println("World!")

	var age int64 = 40

	var flo float64 = 5555.33

	//fmt.Println(age)
	//fmt.Println(flo)

	randNum := 10

	fmt.Println(randNum,age,flo)

	const pi float64 = 3.14159265

	var(
		varA = 2
		varB = 3
	)

	fmt.Println(varA+varB)

	var isOver40 bool = true

	fmt.Printf("%.3f\n",pi)
	fmt.Printf("%d\n",varA)
	fmt.Printf("%c\n",'a')
	fmt.Println(isOver40)

	// For loop
	i:=1

	for(i<=10){
		fmt.Println(i)
		i++
	}

	// For loop and if-else statements
	for j:=0; j<5; j++{

		if j>=3{
			fmt.Println("Greater than or equal to 3")
		} else {
			fmt.Println("Less than 3:(")
		}
	}
	// Switch Statements
	switch varA{
	case 2: fmt.Println("2222222")
	case 4:fmt.Println("4444444")
	default: fmt.Println("Go have fun!")
	}

	//Arrays
	var favNums[5] int

	for i:=0; i<5;i++{
		favNums[i] = i+13;
		fmt.Println(favNums[i])
	}

	// Slices
	numSlice := []int {5,4,3,2,1}
	// takes the 3rd and 4th element
	numSlice2 := numSlice[3:5]
	fmt.Println(numSlice2[1])

	// Going through an array with indexes
	favNums3 := []float64 {1.1,2.2,3.2,4.1,5.1}

	for i,value := range favNums3{
		fmt.Println(i,value)
	}

	// Copying a Slice, appending an element
	mySlice := []int {5,4,3,22,11}
	newSlice := make([]int, 5,10)

	copy(newSlice,mySlice)

	mySlice = append(mySlice,-1)

	fmt.Println("New Index: ",mySlice[5])



	defer fmt.Println("lalala")
	//MAPS
	presAge := make(map[string] int)

	presAge["Berko"] = 42

	fmt.Println("Length of Map: ",len(presAge))

	presAge["Emi"] = 41

	fmt.Println("New Length of Map:",len(presAge))

	fmt.Println("Result of Add Function:",add(favNums3))

	num1,num2 := next2Values(10)
	fmt.Println("Two Return Values:",num1,num2)

	fmt.Println(subtractThem(1,2,3,4,5))

	fmt.Println(factorial(6))

	defer printOne()
	printTwo()

	doubleNum := func(num int) int{
		num=num*2
		return num
	}

	fmt.Println(doubleNum(10))

	fmt.Println(safeDivision(5,0))

	demPanic()

	// Pointers Changing Value
	x := 0
	changeValue(&x)

	fmt.Println("x=",x)
	fmt.Println("Memory Address for x =",&x)




}
// Function Syntax
func add(numbers []float64) float64 {

	sum:=0.0

	for _,value := range numbers{
		sum += value
	}
	return sum
}
// 2 Return Types
func next2Values(number int)(int,int){
	return number+1,number+2
}
// Variatic Function
func subtractThem(args ...int) int{

	finalValue :=0

	for _,value := range args{

		finalValue -= value

	}
	return finalValue
}
// Recursive Factorial
func factorial(num int) int{
	if num==0{
		return 1
	}
	return num*factorial(num-1)
}
// Defer keyword testing
func printOne(){ fmt.Println("I am first")}
func printTwo(){ fmt.Println("I am second")}

// Exception handling
func safeDivision(num1,num2 int) int{

	defer func(){
		recover()
		fmt.Println("ERROR")
	}()

	solution :=num1/num2
	return solution
}
func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()
	panic("PANIC")
}

//Pointers Changing Value
func changeValue(x *int){
	*x = 3
}
