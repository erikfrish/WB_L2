Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]

/*
В го все передается в функции как копия, даже указатели. Слайс является ссылочным типом, он содержит в себе параметры len и cap, а также указатель на массив элементов, поэтому, копируя слайс, вы можете обращаться к отдельным элементам исходного слайса по индексу, потому что скопированный слайс ссылается на тот же массив, что и исходный.
Однако выполняя append мы переаллоцируем массив (кроме случаев, когда cap больше ожидамой len и задан заранее), на который ссылается локальный слайс i. Именно поэтому, изменяя уже элемент по индексу 1 мы меняем его в новом слайсе, который не возвращаяется результатом функции, а значит мы теряем эти изменения после выхода из функции modifySlice.
Чтобы это исправить, достаточно возвращать значение i в функции modifySlice, а результат присваивать исходному слайсу, также как и с append в этой же функции.
*/
```
