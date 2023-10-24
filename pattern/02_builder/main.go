package main

import (
	"02_builder/pattern"
	"fmt"
)

/*
Builder
Giorgi -- это простой грузинский пицца_мастер, он может сделать вам любую пиццу,
которую вы попросите, например пепперони, маргариту и т.д.
Но он не ограничивается пиццей, как и любой профессиональный грузин он умеет делать хачапури,
главное правильно попросить, но даже если вы попросите у него то, чего он не умеет,
он не оставит вас без внимания и угостит хотя бы сигареткой.

Giorgi -- PizzaBuilder, имеет метод Build(what),
достаточно дать ему название и он из него сотворит нужный нам объект нужного типа, следуя своей,
не известной нам, логике, причем не станет возвращать объект, пока не приготовит его полностью

*/

func main() {
	Giorgi := pattern.NewPizzaBuilder("Giorgi")
	pizza := Giorgi.Build("Pepperoni")
	fmt.Printf("\ntype of pizza: %T\npizza: %v\n", pizza, pizza)

	pizza = Giorgi.Build("Adjarian")
	fmt.Printf("\ntype of pizza: %T\npizza: %v\n", pizza, pizza)

	pizza = Giorgi.Build("anything_else?")
	fmt.Printf("\ntype of pizza: %T\npizza: %v\n", pizza, pizza)
}
