package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{10, 12}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
