package main

import (
	"fmt"
	"pro/myutil"
	"time"
)

type person struct {
	name       string
	salary     float64
	age        int
	increments float32
}

func Hello() {
	fmt.Println("Hello World!")
}



func fibanocii(n int) int {
	if n <= 1 {
		return n
	}
	return fibanocii(n-1) + fibanocii(n-2)
}

func inputString() (ans []string, err error) {
	fmt.Println("Enter the number of strings")
	var n int
	fmt.Scanf("%d", &n)
	ans = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter string %d: ", i+1)
		fmt.Scanf("%s", &ans[i])
	}
	return ans, nil
}

func printElement(arr []int, index int) {
	for i := 0; i < index; i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func printing(ab string) (ans string) {
	ans = "Hello  " + ab
	fmt.Println(ans)
	return
}

func print(args any) {
	fmt.Println(args)
}

func MultiplyElement(arr []int, index int) {
	for indexs := 0; indexs < index; indexs++ {
		arr[indexs] *= 2
	}
}
func main() {
	fmt.Println("Values wil multiply here")

	var age int = 123
	var version = "v1"
	fmt.Println(version)

	fmt.Print(age)
	var dimesinon float64 = 4.56
	fmt.Println(dimesinon)

	myutil.Multiply(5, 6)

	personName := "Hewllo bro guys"

	fmt.Println(personName)

	arr := [6]int{1, 2, 3, 4, 5, 6}
	var length = len(arr)

	MultiplyElement(arr[:length], length-1)
	var name string = "Pratap Singh Sisodiya"
	fmt.Println(name)
	printElement(arr[:6], 6)
	fmt.Println(length)
	var ans string = printing("Hello guys")
	fmt.Println(ans)

	const passowrdeky int = 24252325
	fmt.Println(passowrdeky)
	fmt.Print(passowrdeky)

	const variable bool = true
	const floatingVariable float64 = 45.465

	fmt.Println(variable)

	fmt.Printf("%f", floatingVariable)
	numbersArray := [10]int{14, 56, 46, 4, 964, 96846, 846, 54, 6, 4}
	// const variabel = string[]{"pratpa","singh sisodiya"}
	for i := 0; i < 10; i++ {

		fmt.Printf("%d ", numbersArray[i])

	}

	var day string = "no"

	switch day {

	case "Monday", "Tuesday":
		fmt.Println(day)
	case "no":
		fmt.Println("aaj chuthi hian bsdk")
	default:
		fmt.Println("API CHALNE LAGI HAIN")
	}
	var time = time.Now()
	fmt.Println(time)
	var names = &time
	fmt.Println(names)
	var tiems = time.Nanosecond()
	fmt.Println(tiems)
}
