package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// Интерфейс для состояния
type State interface {
	InsertCoin() error
	EjectCoin() error
	SelectDrink() error
	Dispense() error
}

// Конкретные состояния

type NoCoinState struct{}

func (s *NoCoinState) InsertCoin() error {
	fmt.Println("Монета вставлена")
	return nil
}

func (s *NoCoinState) EjectCoin() error {
	fmt.Println("Нет монет для извлечения")
	return nil
}

func (s *NoCoinState) SelectDrink() error {
	fmt.Println("Вставьте монету, чтобы выбрать напиток")
	return nil
}

func (s *NoCoinState) Dispense() error {
	fmt.Println("Вставьте монету и выберите напиток")
	return nil
}

type HasCoinState struct{}

func (s *HasCoinState) InsertCoin() error {
	fmt.Println("Монета уже вставлена")
	return nil
}

func (s *HasCoinState) EjectCoin() error {
	fmt.Println("Монета извлечена")
	return nil
}

func (s *HasCoinState) SelectDrink() error {
	fmt.Println("Напиток выбран")
	return nil
}

func (s *HasCoinState) Dispense() error {
	fmt.Println("Выберите напиток перед выдачей")
	return nil
}

type DispensingState struct{}

func (s *DispensingState) InsertCoin() error {
	fmt.Println("Подождите, выполняется выдача напитка")
	return nil
}

func (s *DispensingState) EjectCoin() error {
	fmt.Println("Невозможно извлечь монету во время выдачи напитка")
	return nil
}

func (s *DispensingState) SelectDrink() error {
	fmt.Println("Напиток уже выбран")
	return nil
}

func (s *DispensingState) Dispense() error {
	fmt.Println("Напиток выдан")
	return nil
}

// Контекст, который хранит текущее состояние
type VendingMachine struct {
	state State
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &NoCoinState{}}
}

// Методы для изменения состояния
func (m *VendingMachine) InsertCoin() {
	m.state.InsertCoin()
	m.setState(&HasCoinState{})
}

func (m *VendingMachine) EjectCoin() {
	m.state.EjectCoin()
	m.setState(&NoCoinState{})
}

func (m *VendingMachine) SelectDrink() {
	m.state.SelectDrink()
	m.setState(&DispensingState{})
}

func (m *VendingMachine) Dispense() {
	m.state.Dispense()
	m.setState(&NoCoinState{})
}

// Метод для установки нового состояния
func (m *VendingMachine) setState(state State) {
	m.state = state
}
