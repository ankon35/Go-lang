package main
import "fmt"


func add(num1 int, num2 int) int {
	sum:= num1 + num2
	return sum
}

func getNumbers(num1 int , num2 int) (int , int){
	sum:= num1 + num2
	mul:= num1 * num2

	return sum , mul
}

func main(){
	var  a int
	var  b int

	fmt.Println("Enter Fisrt Number:")
	fmt.Scan(&a)

	fmt.Println("Enter Second Number:")
	fmt.Scan(&b)

	// sum:= add(a , b)

	// fmt.Printf("The Sum of %d and %d is: %d\n", a , b , sum)

	summ , multi := getNumbers(a,b)
	// mul:= getNumbers(a,b)

	fmt.Printf("The Sum of %d and %d is: %d\n", a , b , summ)
	fmt.Printf("The Mul of %d and %d is: %d\n", a , b , multi)






}