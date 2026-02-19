package main
import "fmt"

func main(){


	var age int

	fmt.Print("Enter your age:")
	fmt.Scan(&age)

	if age >15{
		fmt.Println("You are under 20")
	}else{
		fmt.Println("You are above 20")
	}
		
	

}