package main

// import (
// 	"05_chain_of_resp/pattern"
// )

/*
Chain of responsibility

Для меня идеальной аналогией паттерна цепочка ответственности является госучреждение, где каждый сотрудник посылает тебя к другому сотруднику и кажется, что конца этому не будет.

В данном случае существует интерфейс кабинета с методами Execute(v *Visitor) и SendTo(next Cabinet).
Каждый кабинет может послать посетителя в следующий, пока там не будет выполнена полезная работа. Вызывая в main всего один метод Execute у кабинета 1 мы видим, что посетитель успевает пройти довольно существенный путь, пока не добьется своего. Плюс, если попытаться пройти этот путь повторно его просто развернут.

В реальном коде часто бывает такое, что существует несколько сервисов, которые должны обработать данные последовательно, например middleware авторизации не должен пускать дальше неавторизованных пользователей.
Мы можем построить длинный пайплайн, в котом сервисы будут последовательно вызывать друг друга, а посетитель будет доходить только до положенного ему уровня.
*/
import (
	p "05_chain_of_resp/pattern"
	"fmt"
)

func main() {
	citizen := p.Bear("Alex")
	fmt.Println("Алекс родился и прожил какое-то время нелегально, теперь он хочет получить паспорт.")
	fmt.Println("Для этого он отправляется в кабинет номер 1 местного МВД...")
	p.Cab1.Execute(citizen)
	if citizen.HavePassport {
		fmt.Println("Теперь Алекс может спокойно купить себе пиво, ведь у него есть паспорт.")
		fmt.Println("Или чего покрепче, если начнет доставать паспортный стол:")
		p.Cab1.Execute(citizen)
	}
}
