package main 
import "fmt" 
func main(){ 
	fmt.Println("Enter a number"); 
	var number int
	fmt.Scanln(&number) 
	if(number%2==0){ 
		fmt.Println("Number is even") 
	}
} 