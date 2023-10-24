package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type PizzaBuilder struct {
	building string
}

func NewPizzaBuilder(name string) *PizzaBuilder {
	return &PizzaBuilder{
		building: name,
	}
}

func (pb *PizzaBuilder) Build(building string) any {
	switch building {
	case "Pepperoni":
		return &Pizza{name: building, size_in_cm: 24, ingredients: []string{"Sausages", "Cheese"}}
	case "Margarita":
		return &Pizza{name: building, size_in_cm: 20, ingredients: []string{"Cheese", "Tomato"}}
	case "FourCheeses":
		return &Pizza{name: building, size_in_cm: 35, ingredients: []string{"Cheese1", "Cheese2", "Cheese3", "Cheese4"}}
	case "Carbonara":
		return &Pizza{name: building, size_in_cm: 25, ingredients: []string{"Bacon", "Cheese"}}
	case "Adjarian":
		return &Khachapuri{type_: "Tasty Adjarian Khachapuri"}
	case "Imeruli":
		return &Khachapuri{type_: "Tasty Imeruli Khachapuri"}
	case "Lobiani":
		return &Khachapuri{type_: "Tasty Lobiani Khachapuri"}
	case "Megruli":
		return &Khachapuri{type_: "Tasty Megruli Khachapuri"}
	default:
		return "cigarette"
	}
}

type Pizza struct {
	name        string
	size_in_cm  int
	ingredients []string
}
type Khachapuri struct {
	type_ string
}
