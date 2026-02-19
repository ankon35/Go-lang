package main

import "fmt"

func main(){
	var first_num float64
	var sec_num float64

	fmt.Println("Enter First Number: ")
	fmt.Scan(&first_num)

	fmt.Println("Enter Second Number:")
	fmt.Scan(&sec_num)

	sum := first_num + sec_num

	fmt.Println("The sum of First and Second Number is:",sum)
}