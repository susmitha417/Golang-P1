package methods

import "fmt"

// Defined a struct for Rectangle
type Rectangle struct {
	Width, Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

/*func main() {
    rect := Rectangle{width: 5, height: 10}
    fmt.Println("Rectangle Area:", rect.Area()) // Output: 50
}*/

func (r Rectangle) Display() {
	fmt.Printf("Rectangle [Width: %.2f, Height: %.2f, Area: %.2f]\n", r.Width, r.Height, r.Area())
}
