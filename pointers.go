package main
import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println("Reference ",*p)
	*p = 21
	fmt.Println("Original Declared ",i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
