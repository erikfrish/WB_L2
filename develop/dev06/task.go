package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	options := make(map[string]any)
	options["f"] = flag.String("f", "", `"fields" - выбрать поля (колонки)`)
	options["d"] = flag.String("d", "\t", `"delimiter" - использовать другой разделитель`)
	options["s"] = flag.Bool("s", false, `"separated" - только строки с разделителем`)
	options["o"] = flag.String("o", "\t", `"output-delimiter" - разделяет поля указанным разделителем`)
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("At least 1 filename is required")
		os.Exit(0)
	}
	for _, file_name := range args {
		fmt.Printf("%s:\n\n", file_name)
		res, err := cut(file_name, options)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}

func cut(file_name string, options map[string]any) (string, error) {
	res := ""
	file, err := os.ReadFile(file_name)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(file), "\n")
	fields, err := parseFields(*options["f"].(*string))
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		columns := strings.Split(line, *options["d"].(*string))
		if !*options["s"].(*bool) || len(columns) > 1 {
			subline := make([]string, 0, len(columns))
			if len(fields) == 0 {
				subline = columns
			} else {
				for i, column := range columns {
					if fields[i] {
						subline = append(subline, column)
					}
				}
			}
			res += strings.Join(subline, *options["o"].(*string))
			res += "\n"
		}
	}

	return res, nil
}

func parseFields(f string) (map[int]bool, error) {
	if f == "" {
		return make(map[int]bool, 0), nil
	}
	parts := strings.Split(f, ",")
	res := make(map[int]bool, len(parts))
	for _, part := range parts {
		interval := strings.Split(strings.TrimSpace(part), "-")
		if len(interval) > 1 {
			start, err := strconv.Atoi(interval[0])
			if err != nil {
				return nil, nil
			}
			end, err := strconv.Atoi(interval[len(interval)-1])
			if err != nil {
				return nil, nil
			}
			for num := start; num <= end; num++ {
				res[num] = true
			}
		} else {
			num, err := strconv.Atoi(interval[0])
			if err != nil {
				return nil, nil
			}
			res[num] = true
		}
	}
	return res, nil
}
