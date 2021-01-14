package main
import "fmt"
type Vertex struct {
	name string
	phone int
}
var m = map[string]Vertex{
	"1": Vertex{
		"aditya", 6301819823,
	},
	"2": Vertex{
		"shajulin",9443543746,
	},
}
func main(){
	fmt.Println(m["1"])
	fmt.Println(m["2"])
}
