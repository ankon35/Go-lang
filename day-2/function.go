package main

import "fmt"


func add(num1 int, num2 int){
	sum:= num1 + num2
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
}


func minus(num1 int, num2 int){
	sum:= num1 - num2



	fmt.Printf("The Subtraction of %d and %d is: %d\n" , num1, num2 , sum)
}


func main(){
	var a int 
	var b int

	fmt.Print("Enter first Number: ")
	fmt.Scan(&a)

	fmt.Print("Enter Ssecond Number: ")
	fmt.Scan(&b)

	add(a, b)

	minus(a , b)

	
}