package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
pwd - показать путь до текущего каталога
echo <args> - вывод аргумента в STDOUT
kill <args> - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах
(linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной,
в интерактивном сеансе выводит некое приглашение в STDOUT и ожидает ввода пользователя
через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости
выводит результат на экран.
Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода
(например \quit).
*/

var buf *bytes.Buffer

func main() {
	buf = new(bytes.Buffer)
	sh_loop()
}

func sh_loop() {
	for {
		cur_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		cur_dir_sl := strings.Split(cur_dir, "/")
		if len(cur_dir_sl) > 2 && cur_dir_sl[len(cur_dir_sl)-3] != "/" {
			cur_dir = cur_dir_sl[len(cur_dir_sl)-2] + "/" + cur_dir_sl[len(cur_dir_sl)-1]
		}
		fmt.Printf("[%s]> ", cur_dir)
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pipe := strings.Split(line, "|")
		for i, cmd := range pipe {
			first := (i == 0)
			last := (i == len(pipe)-1)
			args := split_line(cmd)
			cmd_exec(first, last, args...)
		}
	}
}

func cmd_exec(first, last bool, args ...string) error {
	var input io.Reader
	var output io.Writer
	cmd_name := args[0]
	args = args[1:]
	if first {
		input = os.Stdin
	} else {
		input = buf
	}
	if last {
		output = os.Stdout
	} else {
		output = buf
	}
	switch cmd_name {
	case "exit":
		EXIT()
	case "cd":
		return os.Chdir(cmd_name)
	case "echo":
		ECHO(output, args...)
		return nil
	case "help":
		fmt.Println("just use it like poor sh")
	default:
		DEFAULT_CMD(cmd_name, input, output, args...)
	}
	return nil
}

func DEFAULT_CMD(cmd_name string, input io.Reader, output io.Writer, args ...string) error {
	cmd := exec.Command(cmd_name, args...)
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = output
	return cmd.Run()
}

func EXIT() {
	fmt.Println("exiting...")
	os.Exit(0)
}

func ECHO(output io.Writer, args ...string) {
	for _, arg := range args {
		env := strings.TrimLeft(arg, "$")
		if env == arg {
			fmt.Fprint(output, arg, " ")
		} else {
			fmt.Fprintf(output, "%s:\n%s\n", arg, os.Getenv(env))
		}
	}
}

func split_line(line string) []string {
	line = strings.TrimSpace(line)
	args := make([]string, 0, 10)
	word := make([]rune, 0, 10)
	quotes := false
	for _, ru := range line {
		if quotes {
			if ru == '"' {
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
			case '"':
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
