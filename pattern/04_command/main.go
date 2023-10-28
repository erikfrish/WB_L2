package main

/*
Command

Паттерн команда необходим для инкапсуляции какого-то функционала в отдельную сущность
которая вызывается методом Command в данном случае.
Преимуществом создания отдельных объектов команд является отделение
логики пользовательского интерфейса от внутренней бизнес-логики.
Нет нужды разрабатывать отдельные исполнители для каждого вызывающего объекта –
сама команда содержит всю информацию, необходимую для ее исполнения.
Соответственно, ее можно использовать для отсроченного выполнения задачи.
На стороне принимающей команду удобно реализовать очередь на обработку команды, а самой команде
можно добавить метод для отмены и вычеркивание ее из очереди, если она не была выполнена и уже не требуется.
*/

import (
	"04_command/pattern"
	"fmt"
)

func main() {
	tv := pattern.TV{}
	remote := pattern.MakeRemote()
	on := remote.SetButton(pattern.MakeButton(pattern.TurnOnTV{Connected: &tv}))
	off := remote.SetButton(pattern.MakeButton(pattern.TurnOffTV{Connected: &tv}))
	fmt.Printf("Now TV is on: %t\n", tv.IsOn)
	remote.Btns[on].PushBtn()
	fmt.Printf("Now TV is on: %t\n", tv.IsOn)
	remote.Btns[off].PushBtn()
	fmt.Printf("Now TV is on: %t\n", tv.IsOn)
}
