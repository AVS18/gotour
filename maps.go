package main
import "fmt"
type Vertex struct {
	name string
	phone int
}
var m map[string]Vertex
func main() {
	m = make(map[string]Vertex)
	m["1"] = Vertex{
		"aditya",6301819823,
	}
	fmt.Println(m["1"])
}
