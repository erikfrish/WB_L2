package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	options := make(map[string]any)
	options["A"] = flag.Int("A", 0, `"after" печатать +N строк после совпадения`)
	options["B"] = flag.Int("B", 0, `"before" печатать +N строк до совпадения`)
	options["C"] = flag.Int("C", 0, `"context" (A+B) печатать ±N строк вокруг совпадения`)
	options["c"] = flag.Bool("c", false, `"count" (количество строк)`)
	options["i"] = flag.Bool("i", false, `"ignore-case" (игнорировать регистр)`)
	options["v"] = flag.Bool("v", false, `"invert" (вместо совпадения, исключать)`)
	options["F"] = flag.Bool("F", false, `"fixed", точное совпадение со строкой, не паттерн`)
	options["n"] = flag.Bool("n", false, `"line num", печатать номер строки`)
	flag.Parse()
	args := flag.Args()

	for {
		if len(args) < 1 {
			fmt.Println("Type what to search for and name of file to search in: \n(You can use single quote (') to separate prompt with spaces)")
			line, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				args = append(args, split_line(line)...)
			}
		} else if len(args) == 1 {
			fmt.Println("Type name or names of file or files to search in:")
			line, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				args = append(args, split_line(line)...)
			}
		} else {
			break
		}
	}

	prompt := args[0]
	for _, file_name := range args[1:] {
		if err := search_in_file(prompt, file_name, options); err != nil {
			fmt.Println(err)
		}
	}

}

func search_in_file(prompt, file_name string, options map[string]any) error {
	var indexes = make([]int, 0)
	bytes, err := os.ReadFile(file_name)
	if err != nil {
		return err
	}
	if *options["i"].(*bool) {
		prompt = strings.ToLower(prompt)
	}
	data := strings.Split(string(bytes), "\n")
	for i, str := range data {
		if *options["i"].(*bool) {
			str = strings.ToLower(str)
		}
		if *options["F"].(*bool) {
			if strings.TrimSpace(str) == strings.TrimSpace(prompt) {
				indexes = append(indexes, i)
			}
		} else if strings.Contains(str, prompt) {
			indexes = append(indexes, i)
		}
	}
	if len(indexes) == 0 {
		return fmt.Errorf("Nothing found")
	}

	print_result(data, indexes, options)
	return nil
}

func print_result(data []string, indexes []int, options map[string]any) {
	to_print := ""
	before := *options["B"].(*int)
	after := *options["A"].(*int)
	if *options["C"].(*int) != 0 {
		before, after = *options["C"].(*int), *options["C"].(*int)
	}
	if *options["v"].(*bool) {
		indexes_map := make(map[int]bool)
		for _, i := range indexes {
			indexes_map[i] = true
		}
		if *options["c"].(*bool) {
			to_print = fmt.Sprintf("Found %d lines:\n", len(data)-len(indexes))
		}
		for i, line := range data {
			if !indexes_map[i] {
				if *options["n"].(*bool) {
					to_print += fmt.Sprintf("%d ", i)
				}
				to_print += fmt.Sprintf("%s\n", line)
			}
		}
	} else {
		if *options["c"].(*bool) {
			to_print = fmt.Sprintf("Found %d lines:\n", len(indexes))
		}
		for _, index := range indexes {
			for i := index - before; i <= index+after; i++ {
				if *options["n"].(*bool) {
					to_print += fmt.Sprintf("%d ", i)
				}
				to_print += fmt.Sprintf("%s\n", data[i])
			}
		}
	}

	fmt.Println(to_print)
}

func split_line(line string) []string {
	line = strings.TrimSpace(line)
	args := make([]string, 0, 10)
	word := make([]rune, 0, 10)
	quotes := false
	for _, ru := range line {
		if quotes {
			if ru == '\'' {
				quotes = false
				if len(word) > 0 {
					args = append(args, string(word))
					word = make([]rune, 0, 10)
				}
			} else {
				word = append(word, ru)
			}
		} else {
			switch ru {
			case ' ':
				if len(word) > 0 {
					args = append(args, string(word))
					word = make([]rune, 0, 10)
				}
			case '\'':
				quotes = true
			default:
				word = append(word, ru)
			}
		}
	}
	if len(word) > 0 {
		args = append(args, string(word))
	}
	return args
}
