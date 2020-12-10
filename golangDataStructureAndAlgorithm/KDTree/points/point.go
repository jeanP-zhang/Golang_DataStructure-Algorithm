package points

import "fmt"

type Points interface {
	Dimensions() int         //尺寸
	Dimension(i int) float64 //维度
}

type Point struct {
	Coordinates []float64
	Data        interface{}
}

//NewPoint creats a new point at the given coordinates and contains the given data
func NewPoint(coordinates []float64, data interface{}) *Point {
	return &Point{
		coordinates,
		data,
	}
}

//Dimensions returns the total number of dimensions
func (p *Point) Dimensions() int {
	return len(p.Coordinates)
}

//Dimension returns the value of the i-th dimension
func (p *Point) Dimension(i int) float64 {
	return p.Coordinates[i]
}

//Strings returns the string representation of the point
func (p *Point) Strings() string {
	return fmt.Sprintf("{%v,%v}", p.Coordinates, p.Data)
}
