package main
import "fmt"

func main(){
	secret:= 10

	var guess int 

	var press string

	for{
		fmt.Print("Enter your guess:")
		fmt.Scan(&guess)

		if guess==secret{
			fmt.Print("YOu are correct")
			break
		}else{
			fmt.Println("You are wrong. Do you want to continue press A for continue")

			fmt.Scan(&press)

			if press!="A"{
				break
			}

		}
	}
}