package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

/*
=== Базовое решение ===

Создан пакет what_time, в котором используется библиотека github.com/beevik/ntp
она импортируется в таск и используется для получения точного времени.
В случае ненулевой ошибки программа завершается с кодом ошибки 1, перед этим выводя
текст ошибки в stderr.
У пакета what_time есть метод для получения времени в формате
time.Time, но также есть вариант сразу отформатировать результат в строку с помощью time.Format
*/
import (
	"fmt"
	"os"
	"time"

	wt "dev01/what_time"
)

func main() {
	ntp_server := "3.beevik-ntp.pool.ntp.org"
	ntp := wt.New(ntp_server)
	now, err := ntp.GetTimeF(time.StampMilli)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	fmt.Println(now)
	os.Exit(0)
}
