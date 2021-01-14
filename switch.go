package main
import ("fmt" 
"runtime")
func main(){
	var x string;
	fmt.Println("Enter a programming language")
	fmt.Scanln(&x);
	switch(x){
		case "c":
			fmt.Println("C Language") 
		case "cpp":
			fmt.Println("C++ Language")
		case "java":
			fmt.Println("Java Language")
	}
	fmt.Println(runtime.GOOS)
}