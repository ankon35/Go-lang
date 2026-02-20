package main
import "fmt"

func welcome_msg(){
	fmt.Print("Welcome to the system\n")
}

func getUser()(string , string){

	var name string
	var email string
	
	fmt.Printf("Enter you name: ")
	fmt.Scan(&name)

	fmt.Printf("Enter you email: ")
	fmt.Scan(&email)

	return name , email

}


func thanks_msg(name string){
	fmt.Print("Thank You ", name)
}


func main(){
	welcome_msg()

	name1, email1:= getUser()

	fmt.Println("Name:", name1)
	fmt.Println("Email:", email1)

	thanks_msg(name1)
}
