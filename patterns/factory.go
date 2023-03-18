package patterns

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("circle")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("rectangle")
}

type ShapeFactory struct{}

func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "circle" {
		return &Circle{}
	} else if shapeType == "rectangle" {
		return &Rectangle{}
	} else {
		return nil
	}
}

func main() {
	factory := &ShapeFactory{}

	circle := factory.GetShape("circle")
	circle.Draw()

	rectangle := factory.GetShape("rectangle")
	rectangle.Draw()
}
