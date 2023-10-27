package pattern

import "math"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Square struct {
	a float64
}

func MakeSquare(a float64) *Square {
	return &Square{a: a}
}
func (s *Square) GetType() string {
	return "square"
}
func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

type Rectangle struct {
	a, b float64
}

func MakeRectangle(a, b float64) *Rectangle {
	return &Rectangle{a: a, b: b}
}

func (r *Rectangle) GetType() string {
	return "rectangle"
}
func (r *Rectangle) Accept(v Visitor) {
	v.visitForRectangle(r)
}

type Circle struct {
	rad float64
}

func MakeCircle(rad float64) *Circle {
	return &Circle{rad: rad}
}

func (c *Circle) GetType() string {
	return "circle"
}
func (c *Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}

type CalculateArea struct {
	Area float64
}

func (ca *CalculateArea) visitForSquare(s *Square) {
	ca.Area = s.a * s.a
}
func (ca *CalculateArea) visitForRectangle(r *Rectangle) {
	ca.Area = r.a * r.b
}
func (ca *CalculateArea) visitForCircle(c *Circle) {
	ca.Area = math.Pi * c.rad * c.rad
}
