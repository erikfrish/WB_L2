package main

import (
	"01_facade/pattern"
	"fmt"
)

/*
Facade

в данном примере мы создаем фасад PizzaFacade, который оперирует сразу 2 сущностями -- кошельком и пиццерией.
так как оперировать ими по отдельности было бы сложно для Джека и Люси, они создают фасад, которому передают свои кошельки
и пиццерию, в которой хотели бы сделать заказ, потом спрашивают "Могу ли я заказать такую пиццу?" у фасада, а он
проверяет состояние их кошельков и наличие пиццы в пиццерие, выдавая четкий ответ.
У фасада реализован только один метод CanIOrderAPizzaNow(pizzaName), но в теории можно добавить методы для оформления
заказа или отмены, Джеку и Люси будет необходимо лишь обратиться к своему фасаду с просьбой это сделать без необходимости
самостоятельно связываться с пиццерией и вынимать деньги из кошельков, за них это сделает фасад,
представим, что он их дворецкий, которому они доверили свои кошельки и послали в пиццерию
*/

func main() {
	// making brown wallet for Jack with $100 and green for Lucy with $1000
	JackWallet := pattern.MakeWallet("Jack", "brown", 100)
	LucyWallet := pattern.MakeWallet("Lucy", "green", 1000)
	pizzeria := pattern.MakePizzeria(9, 21) // making Pizzeria that works from 9:00 to 21:00
	// adding Pepperoni and MegaPepperoni to menu
	pizzeria.AddToMenu(pattern.MakePizza("Pepperoni", 10))
	pizzeria.AddToMenu(pattern.MakePizza("MegaPepperoni", 150))
	// creating facade to work with certain pizzeria and your certain wallet
	JackF := pattern.MakePizzaFacade(JackWallet, pizzeria)
	LucyF := pattern.MakePizzaFacade(LucyWallet, pizzeria)
	// Jack can order a Pepperoni with his cash, but can't order MegaPepperoni
	// Lucy can order both of them, but can't order a pizza not from menu
	fmt.Println(JackF.CanIOrderAPizzaNow("Pepperoni"))
	fmt.Println(JackF.CanIOrderAPizzaNow("MegaPepperoni"))
	fmt.Println(LucyF.CanIOrderAPizzaNow("Pepperoni"))
	fmt.Println(LucyF.CanIOrderAPizzaNow("MegaPepperoni"))
	fmt.Println(LucyF.CanIOrderAPizzaNow("FancyPizza"))
}
