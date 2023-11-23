package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type IGun interface {
	Shoot()
	setName(name string)
	setDamage(damage int)
	GetName() string
	GetDamage() int
}

type Gun struct {
	name   string
	damage int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) setDamage(Damage int) {
	g.damage = Damage
}

func (g *Gun) GetDamage() int {
	return g.damage
}
func (g *Gun) Shoot() {
	fmt.Printf("Bang! You gained %d of damage\n", g.damage)
}

type Laser struct {
	Gun
}

func newLaser() IGun {
	return &Laser{
		Gun: Gun{
			name:   "Laser gun",
			damage: 4,
		},
	}
}

func newMiniGun() IGun {
	return &Laser{
		Gun: Gun{
			name:   "Mini Gun",
			damage: 2,
		},
	}
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "laser" {
		return newLaser(), nil
	}
	if gunType == "minigun" {
		return newMiniGun(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
