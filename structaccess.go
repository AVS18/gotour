package main
import "fmt"
type node struct{
	data int
	fdata float64
}
func main(){	
	d,f := 0,0.0
	fmt.Println("Enter the integer data and Float data")
	fmt.Scanln(&d)
	fmt.Scanln(&f)
	t := node{d,f}
	fmt.Println("From strucutre ",t.data,t.fdata)
}