package sorting

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Months = map[string]int{
	"january":   0,
	"february":  1,
	"march":     2,
	"april":     3,
	"may":       4,
	"june":      5,
	"july":      6,
	"august":    7,
	"september": 8,
	"october":   9,
	"november":  10,
	"december":  11,
}

var Multipliers = map[string]float64{
	"B": 1,
	"K": 1024,
	"M": 1024 * 1024,
	"G": 1024 * 1024 * 1024,
	"T": 1024 * 1024 * 1024 * 1024,
	"P": 1024 * 1024 * 1024 * 1024 * 1024,
}

func isMonth(month string) bool {
	for m := range Months {
		if m == strings.ToLower(month) {
			return true
		}
	}
	return false
}

func isNumWithSuffix(str string) bool {
	regex := regexp.MustCompile(`^(\d+\.*\d*[B|K|M|G|T|P]{1})$`)
	return regex.MatchString(str)
}

func DefaultSort(units map[string][]string) []string {
	sorted := make([]string, 0)
	keys := make([]string, 0)
	for key := range units {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		sorted = append(sorted, units[key]...)
	}
	return sorted
}

func NumSort(units map[string][]string) []string {
	sorted := make([]string, 0)
	numKeys := make([]int, 0)
	strKeys := make([]string, 0)
	for key := range units {
		if num, err := strconv.Atoi(key); err == nil {
			numKeys = append(numKeys, num)
		} else {
			strKeys = append(strKeys, key)
		}
	}
	sort.Ints(numKeys)
	sort.Strings(strKeys)
	for _, k := range strKeys {
		sorted = append(sorted, units[k]...)
	}
	for _, k := range numKeys {
		sorted = append(sorted, units[strconv.Itoa(k)]...)
	}
	return sorted
}

func MonthSort(units map[string][]string) []string {
	sorted := make([]string, 0)
	monthKeys := make([]string, 0)
	otherKeys := make([]string, 0)

	for key := range units {
		if isMonth(key) {
			monthKeys = append(monthKeys, key)
		} else {
			otherKeys = append(otherKeys, key)
		}
	}
	sort.Slice(monthKeys, func(i, j int) bool {
		m1 := strings.ToLower(monthKeys[i])
		m2 := strings.ToLower(monthKeys[j])
		return Months[m1] < Months[m2]
	})
	sort.Strings(otherKeys)
	for _, k := range otherKeys {
		sorted = append(sorted, units[k]...)
	}
	for _, k := range monthKeys {
		sorted = append(sorted, units[k]...)
	}
	return sorted
}

func SuffixSort(units map[string][]string) []string {
	sorted := make([]string, 0)
	suffixKeys := make([]string, 0)
	otherKeys := make([]string, 0)

	for key := range units {
		if isNumWithSuffix(key) {
			suffixKeys = append(suffixKeys, key)
		} else {
			otherKeys = append(otherKeys, key)
		}
	}
	sort.Slice(suffixKeys, func(i, j int) bool {
		v1 := strings.ToLower(suffixKeys[i])
		v2 := strings.ToLower(suffixKeys[j])
		return NumWithSuffixToRealNum(v1) < NumWithSuffixToRealNum(v2)
	})
	sort.Strings(otherKeys)
	for _, k := range otherKeys {
		sorted = append(sorted, units[k]...)
	}
	for _, k := range suffixKeys {
		sorted = append(sorted, units[k]...)
	}
	return sorted
}

func NumWithSuffixToRealNum(str string) float64 {
	num, err := strconv.ParseFloat(str[:len(str)-1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	suf := strings.ToUpper(string(str[len(str)-1]))
	return num * Multipliers[suf]
}

type Sorter struct {
	options map[string]*bool
	k       *int
}

func NewSorter(options map[string]*bool, k *int) *Sorter {
	return &Sorter{options: options, k: k}
}

func (s *Sorter) Run(filename string, toSave *string) {
	// читаем весь файл в слайс строк
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("There is no such file", err)
		os.Exit(1)
	}
	text := strings.Split(string(bytes), "\n")
	var sorted []string

	// убираем пробелы
	if *s.options["b"] {
		for i := range text {
			text[i] = strings.TrimSpace(text[i])
		}
	}
	// убираем повторения
	if *s.options["u"] {
		ma := make(map[string]struct{}, len(text))
		for _, str := range text {
			ma[str] = struct{}{}
		}
		text = make([]string, 0, len(text))
		for str := range ma {
			text = append(text, str)
		}
	}

	// подготавливаем текст для сортировки, ищем ключи и формируем мапу по ключам со слайсами строк
	units := groupBySortKeys(text, *s.k)

	// сортируем по выбранному ключу (стоит выбирать только один из них)
	if *s.options["n"] {
		sorted = NumSort(units)
	} else if *s.options["h"] {
		sorted = SuffixSort(units)
	} else if *s.options["M"] {
		sorted = MonthSort(units)
	} else {
		sorted = DefaultSort(units)
	}

	// переворачиваем результат
	if *s.options["r"] {
		for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
			sorted[i], sorted[j] = sorted[j], sorted[i]
		}
	}
	fmt.Println("\n", text, len(text))

	if *s.options["c"] {
		if Equal(text, sorted) {
			fmt.Println("File is already sorted")
		} else {
			fmt.Println("File is not sorted yet")
		}
		return
	} else {
		res_file, err := os.Create(*toSave)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res_bytes := make([]byte, 0, len(bytes))
		for _, v := range sorted {
			res_bytes = append(res_bytes, v...)
			res_bytes = append(res_bytes, byte('\n'))
		}
		// fmt.Println(string(res_bytes))
		_, err = res_file.Write(res_bytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func groupBySortKeys(toSort []string, k int) map[string][]string {
	units := make(map[string][]string, len(toSort))
	for _, str := range toSort {
		columns := strings.Split(str, " ")
		sortColumn := columns[0]
		if k < len(columns) {
			sortColumn = columns[k]
		}
		sortColumn = strings.TrimSpace(sortColumn)
		units[sortColumn] = append(units[sortColumn], str)
	}
	return units
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
