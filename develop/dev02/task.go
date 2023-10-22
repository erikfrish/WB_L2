package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
=== Решение на распаковку ===

Для распаковки используется функция Unpack она оперирует строкой как слайсом рун
сперва валидируя строку с помощью validateString(), затем распаковывает.
Каждый символ проверяется на один из трех вариантов: он цифра, тогда записываем в number
(нельзя сразу делать преобразование, так как в строке может быть число из нескольких цифр),
он escapedRune (\), тогда просто записываем следующий символ в итоговую строку, и остальные случаи,
тогда либо записываем букву в строку, либо записываем ее столько раз, сколько указано в number
*/

import (
	"fmt"
	"strconv"
	"unicode"
)

const escapedRune = 92

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack(`qwe\4\5`))
	fmt.Println(Unpack("45"))
	fmt.Println(Unpack(""))
}

func Unpack(str string) (string, error) {
	var res, number []rune
	runes := []rune(str)

	if err := validateString(runes); err != nil {
		return "", err
	}

	var previous rune
	for i := 0; i < len(runes); i++ {
		rune := runes[i]
		if unicode.IsDigit(rune) {
			number = append(number, rune)
		} else if rune == escapedRune {
			previous = runes[i+1]
			res = append(res, runes[i+1])
			i++
		} else {
			if len(number) != 0 {
				num, err := strconv.Atoi(string(number))
				if err != nil {
					return "", err
				}
				if num == 0 {
					res = res[:len(res)-1]
				}
				for j := 0; j < num; j++ {
					res = append(res, previous)
				}
				number = nil
			}
			res = append(res, rune)
			previous = rune
		}
	}

	return string(res), nil
}

func validateString(r []rune) error {
	// Первый символ не должен быть цифрой, а последний не должен быть escapedRune (\)
	if len(r) > 0 {
		if unicode.IsDigit(r[0]) {
			return fmt.Errorf("error: string starts with a number")
		} else if r[len(r)-1] == escapedRune {
			return fmt.Errorf("error: last symbol is escaped symbol")
		}
	}
	return nil
}
