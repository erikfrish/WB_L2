Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

/*
В выводе мы видим error, потому что переменная err содержит в себе nil, но не эквивалентна nil. Также как и в предыдущем вопросе мы объявляем переменную интерфейсного типа error, а потом определяем ее тип, но записываем ниловое значение. Если вывести значение err, мы увидим nil, но проверка на nil будет false, так как в err хранится тип *customError.
Можно исправить это с помощью type assertion, достаточно сравнивать с nil не err, а err.(*customError), тогда в выводе мы увидим ok

*/
```
