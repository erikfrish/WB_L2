package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

var (
	Cab1          = Cabinet1{}
	cab3          = Cabinet3{}
	cab53         = VaccineCabinet{}
	taxPayCabinet = cashRegister{}
	cab32         = registrationCabinet{}
)

type Cabinet interface {
	Execute(v *Visitor)
	SendTo(next Cabinet)
}

type Visitor struct {
	name             string
	vaccineDone      bool
	taxPayed         bool
	registrationDone bool
	HavePassport     bool
}

func Bear(name string) *Visitor {
	return &Visitor{name: name}
}

type Cabinet1 struct {
	next Cabinet
}

func (c *Cabinet1) Execute(v *Visitor) {
	fmt.Printf("Здравствуйте, %s, тут вы можете получить новый паспорт,	но мы обслуживаем только привитых, тех, кто оплатил пошлину и имеет постоянную регистрацию.\n", v.name)
	if v.HavePassport {
		fmt.Printf("Погоди, так у тебя уже есть один, мы тут благотворительностю не занимаемся, пшел отсюда!\n\n")
	} else if !v.vaccineDone {
		fmt.Printf("Не вижу у вас QR-кода о прививке, приходите, когда он будет. Сделать можно в кабинете 53\n\n")
		c.SendTo(&cab53)
		c.next.Execute(v)
	} else if !v.taxPayed {
		fmt.Printf("А вы пошлину уже оплатили? Где квитанция? То-то же, вам в кассу.\n\n")
		c.SendTo(&taxPayCabinet)
		c.next.Execute(v)
	} else if !v.registrationDone {
		fmt.Printf("Молодцы, пошлину оплатили, но я не выдам вам документ без регистрации, хотя бы временной. Идите разбирайтесь в 32-ой.\n\n")
		c.SendTo(&cab32)
		c.next.Execute(v)
	} else {
		fmt.Printf("Поздравляю, вы на финишной прямой, мы займемся вашим документом. Когда-нибудь. Приходите через 3 месяца\n\n")
		for month := 0; month < 3; month++ {
			time.Sleep(time.Second)
			fmt.Printf("Прошел %d месяц...\n", month+1)
		}
		fmt.Printf("О, это снова вы! Да, можете забирать в третьем кабинете.\n\n")
		c.SendTo(&cab3)
		c.next.Execute(v)
	}
}
func (c *Cabinet1) SendTo(next Cabinet) {
	c.next = next
}

type VaccineCabinet struct {
	next Cabinet
}

func (c *VaccineCabinet) Execute(v *Visitor) {
	fmt.Printf("Здравствуйте! Раньше чернухой болели? Твиттер лечили? Это я так, из вежливости, на самом деле мне пофиг. Подставляй плечо, %s.\n\n", v.name)
	time.Sleep(time.Second * 2)
	v.vaccineDone = true
	fmt.Printf("Поздравляю, QR-код получите на госуслугах, до свидания.\n\n")
	c.SendTo(&Cab1)
	c.next.Execute(v)
}

func (c *VaccineCabinet) SendTo(next Cabinet) {
	c.next = next
}

type Cabinet3 struct {
	next Cabinet
}

func (c *Cabinet3) Execute(v *Visitor) {
	fmt.Printf("Привет, %s! Вам все понравилось? Оставьте пожалуйста отзыв о получении услуг, мы будем очень благодарны!\n\n", v.name)
	time.Sleep(time.Second * 2)
	v.HavePassport = true
	fmt.Printf("Вот ваш паспорт, приходите еще <3\n\n")
}
func (c *Cabinet3) SendTo(next Cabinet) {
	c.next = next
}

type cashRegister struct {
	next Cabinet
}

func (c *cashRegister) Execute(v *Visitor) {
	fmt.Printf("Здравствуйте, %s!, с вас 100500 рублей. Платить сюда.\n\n", v.name)
	time.Sleep(time.Second * 2)
	fmt.Printf("*Идет оплата 100500 рублей*\n")
	time.Sleep(time.Second * 2)
	v.taxPayed = true
	fmt.Printf("Вот ваш чек, без него вы, ну знаете, без бумажки ты кто. До свидания.\n\n")
	c.SendTo(&Cab1)
	c.next.Execute(v)
}
func (c *cashRegister) SendTo(next Cabinet) {
	c.next = next
}

type registrationCabinet struct {
	next Cabinet
}

func (c *registrationCabinet) Execute(v *Visitor) {
	fmt.Printf("Не добрый день, где живете? О, так соседи будем. Если что, я вас не знаю, вы меня не знаете.\n\n")
	time.Sleep(time.Second * 2)
	v.registrationDone = true
	fmt.Printf("Вот ваша регистрация, на бумажке конечно, возвращайтесь после получения паспорта, нормальную поставим. А лучше не возвращайтесь.\n\n")
	c.SendTo(&Cab1)
	c.next.Execute(v)
}
func (c *registrationCabinet) SendTo(next Cabinet) {
	c.next = next
}
