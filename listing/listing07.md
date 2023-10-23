Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```go
1
2
3
4
6
5
8
7
// порядок может быть другим 
0
0
0
...
0
... и так до бесконечности

Дело в том, что каналы a и b закрыты, но в функции merge мы все равно продолжаем читать из них данные. В этом случае, читая только данные мы будем получать нулевое значение для типа, хранящегося в закрытом канале. Позже мы читаем из канала c эти нулевые данные. Чтобы этого избежать следует добавить в кейсы селекта проверку на закрытость обоих каналов, тогда после того, как a и b закроются, закроется и c, а в main мы перестанем итерироваться по range. Вот вариант исправленой функции merge:

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		outerloop:
			for {
				select {
				case v, ok := <-a:
					if !ok {
						_, ok2 := <-b
						if !ok2 {
							break outerloop
						}
					}
				c <- v
				case v, ok := <-b:
					if !ok {
					_, ok1 := <-a
						if !ok1 {
							break outerloop
						}
					}
				c <- v
				}
			}
		close(c)
	}()
	return c
}
```
