package main

import (
	"fmt"
)

/*
3. Interfaces: ระบบคำนวณพื้นที่ (Shape)
โจทย์: ให้ประกาศ (Define) Interface ชื่อ Shape ที่มี method ชื่อ Area() float64
สร้าง Struct 2 ตัวชื่อ Rectangle (มี field Width, Height) และ Circle (มี field Radius)
Implement method Area ให้กับ Struct ทั้งสองเพื่อคำนวณพื้นที่
เขียนฟังก์ชันแยกออกมาน 1 ตัว ชื่อ PrintArea(s Shape) ที่รับ Shape รูปทรงไหนก็ได้แล้วพิมพ์ นำดพื้นที่ออกมาน
*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	PrintArea(rect)
	PrintArea(circle)
}

func PrintArea(s Shape) {
	switch shape := s.(type) {
	case Rectangle:
		fmt.Printf("Area of Rectangle (%.2f x %.2f): %.2f\n", shape.Width, shape.Height, shape.Area())
	case Circle:
		fmt.Printf("Area of Circle (radius %.2f): %.2f\n", shape.Radius, shape.Area())
	default:
		fmt.Println("Unknown shape")
	}
}
