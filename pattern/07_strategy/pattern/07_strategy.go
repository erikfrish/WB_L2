package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретная стратегия 1
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с помощью кредитной карты", amount)
}

// Конкретная стратегия 2
type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f через PayPal", amount)
}

// Контекст, использующий стратегию
type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
}

// Метод для установки стратегии
func (cart *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	cart.PaymentStrategy = strategy
}

// Метод для проведения платежа с использованием текущей стратегии
func (cart *ShoppingCart) ProcessPayment(amount float64) string {
	return cart.PaymentStrategy.Pay(amount)
}
