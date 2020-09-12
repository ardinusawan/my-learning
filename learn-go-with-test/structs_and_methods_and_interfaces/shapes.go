package main

// Rectangle element
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter take width and height and return calculated the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area take width and height and return area of rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
