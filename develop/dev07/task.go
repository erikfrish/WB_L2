package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

import (
	"fmt"
	"time"
)

func main() {

	done1 := make(chan any, 0)
	done2 := make(chan any, 0)
	done3 := make(chan any, 0)
	done := UniteDoneChannels(done1, done2, done3)

	go func(ch chan any) {
		fmt.Println("routine1")
		time.Sleep(time.Second * 4)
		ch <- "done1"
	}(done1)

	go func(ch chan any) {
		fmt.Println("routine2")
		time.Sleep(time.Second * 50)
		ch <- "done2"
	}(done2)

	go func(ch chan any) {
		fmt.Println("routine3")
		time.Sleep(time.Second * 2)
		ch <- "done3"
	}(done3)

	fmt.Println(<-done)
}

func UniteDoneChannels(chans ...chan any) chan any {
	res := make(chan any, 0)
	for _, ch := range chans {
		go func(ch <-chan interface{}) {
			for val := range ch {
				res <- val
			}
		}(ch)
	}
	return res
}
