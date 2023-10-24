package pattern

// "github.com/golang-module/carbon"

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"strconv"
	"time"
)

// Описание самого фасада
type PizzaFacade struct {
	wallet   *Wallet
	pizzeria *Pizzeria
}

func MakePizzaFacade(w *Wallet, p *Pizzeria) *PizzaFacade {
	return &PizzaFacade{
		wallet:   w,
		pizzeria: p,
	}
}

// Метод фасада, позволяющий проверить доступность пиццы для заказа
func (pf *PizzaFacade) CanIOrderAPizzaNow(wanted string) string {
	answer := ""
	wantedP := pf.pizzeria.GetMenu()[wanted]
	if (wantedP != Pizza{}) &&
		(pf.wallet.amount > wantedP.price) &&
		(pf.pizzeria.IsWorking()) {
		answer = "Yes, " + pf.wallet.owner + ", U can buy " + wantedP.name + " by paying $" + strconv.Itoa(wantedP.price) +
			" from your " + pf.wallet.color + " wallet"
	} else {
		answer = "Sorry, " + pf.wallet.owner + ", it's not your day =-("
	}
	return answer
}

type Wallet struct {
	owner  string
	color  string
	amount int
}

func MakeWallet(o, c string, a int) *Wallet {
	return &Wallet{
		owner:  o,
		color:  c,
		amount: a,
	}
}

type Pizza struct {
	name  string
	price int
}

func MakePizza(n string, p int) Pizza {
	return Pizza{
		name:  n,
		price: p,
	}
}

type Pizzeria struct {
	menu           map[string]Pizza
	workHoursStart int
	workHoursEnd   int
}

func MakePizzeria(whs, whe int) *Pizzeria {
	return &Pizzeria{
		menu:           map[string]Pizza{},
		workHoursStart: whs,
		workHoursEnd:   whe,
	}
}

func (p *Pizzeria) GetMenu() map[string]Pizza {
	return p.menu
}

func (p *Pizzeria) AddToMenu(pi Pizza) {
	p.menu[pi.name] = pi
}

func (p *Pizzeria) IsWorking() bool {
	now := time.Now().Hour()
	if now >= p.workHoursStart && now < p.workHoursEnd {
		return true
	}
	return false
}
