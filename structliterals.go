package main
import "fmt"
type Vertex struct {
	var1, var2 int
}
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{var1: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)
func main() {
	fmt.Println(v1,v2,v3,p)
}
