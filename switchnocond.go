package main
import (
	"fmt"
	"time"
)
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Attend Online Class")
	case t.Hour() < 17:
		fmt.Println("Attend Lab")
	default:
		fmt.Println("Be Productive Today")
	}
}
