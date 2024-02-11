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

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	r, err := ntp.Query("0.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to request time data: %s\n", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	exactTime := currentTime.Add(r.ClockOffset)

	fmt.Printf("Текущее время: %s\n", currentTime.Format(time.UnixDate))
	fmt.Printf("Точное время: %s\n", exactTime.Format(time.UnixDate))
}
