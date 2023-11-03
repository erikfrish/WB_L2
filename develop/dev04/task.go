package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	dict1 := []string{"пятак", "пятка", "тяпка"}
	fmt.Println(findAnagrams(&dict1))
	dict2 := []string{"листок", "слиток", "столик"}
	fmt.Println(findAnagrams(&dict2))
	dict3 := append(dict1, dict2...)
	res3 := findAnagrams(&dict3)
	for k, v := range *res3 {
		fmt.Println("key =", k)
		for _, v2 := range v {
			fmt.Println("el =", v2)
		}
	}
}

func findAnagrams(dict *[]string) *map[string][]string {
	anagrams := make(map[string][]string, 0)
	for _, v := range *dict {
		word := strings.ToLower(v)
		runes := []rune(word)
		strs := make([]string, 0, len(runes))
		for _, rune := range runes {
			strs = append(strs, string(rune))
		}
		sort.Strings(strs)
		key := ""
		for _, s := range strs {
			key += s
		}
		anagrams[key] = append(anagrams[key], word)
	}
	res := make(map[string][]string, 0)
	for _, words := range anagrams {
		key := words[0]
		temp_m := make(map[string]struct{}, len(words))
		for _, word := range words {
			temp_m[word] = struct{}{}
		}
		words = make([]string, 0, len(temp_m))
		for word := range temp_m {
			words = append(words, word)
		}
		if len(words) < 2 {
			continue
		}
		sort.Strings(words)
		res[key] = words
	}
	return &res
}
