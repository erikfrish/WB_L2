package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Command()
}

type Remote struct {
	Btns []*Button
}

func MakeRemote() *Remote {
	return &Remote{}
}
func (r *Remote) SetButton(b *Button) int {
	r.Btns = append(r.Btns, b)
	return len(r.Btns) - 1
}

type Button struct {
	cmd Command
}

func MakeButton(cmd Command) *Button {
	return &Button{cmd: cmd}
}
func (b *Button) PushBtn() {
	b.cmd.Command()
}
func (b *Button) SetBtnCommand(cmd Command) {
	b.cmd = cmd
}

type TV struct {
	IsOn bool
}

func (t *TV) on() {
	t.IsOn = true
	fmt.Println("Turning TV on")
}

func (t *TV) off() {
	t.IsOn = false
	fmt.Println("Turning TV off")
}

type TurnOnTV struct {
	Connected *TV
}
type TurnOffTV struct {
	Connected *TV
}

func (cmd TurnOnTV) Command() {
	cmd.Connected.on()
}
func (cmd TurnOffTV) Command() {
	cmd.Connected.off()
}
