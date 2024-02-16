package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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

// builder добавляет к strings.Builder метод writeRunes
type builder struct {
	strings.Builder
}

// writeRunes записывает в builder n рун c
func (b *builder) writeRunes(c rune, n int) {
	b.Grow(n)

	for i := 0; i < n; i++ {
		b.WriteRune(c)
	}
}

func main() {
	v, err := unpackString("a4bc2d5e")
	fmt.Printf("%q %v\n", v, err)

	v, err = unpackString("abcd")
	fmt.Printf("%q %v\n", v, err)

	v, err = unpackString("45")
	fmt.Printf("%q %v\n", v, err)

	v, err = unpackString("")
	fmt.Printf("%q %v\n", v, err)
}

func unpackString(str string) (string, error) {
	var prevChar rune
	var cnt int

	var b builder

	for _, c := range str {
		// Если цифра, увеличение счётчика
		if unicode.IsDigit(c) {
			if prevChar == 0 {
				return "", errors.New("(некорректная строка)")
			}

			err := increaseCounter(&cnt, c)
			if err != nil {
				return "", nil
			}

			continue
		}

		if prevChar != 0 {
			if cnt == 0 {
				cnt++ // чтобы записывать символы без числа подсле них
			}

			b.writeRunes(prevChar, cnt)
			cnt = 0
		}

		prevChar = c
	}

	// Обработка последнего символа
	if prevChar != 0 {
		if unicode.IsDigit(prevChar) {
			err := increaseCounter(&cnt, prevChar)
			if err != nil {
				return "", err
			}
		}

		if cnt == 0 {
			cnt++ // чтобы записывать символы без числа подсле них
		}

		b.writeRunes(prevChar, cnt)
	}

	return b.String(), nil
}

// increaseCounter добавляет цифру c в конец числа cnt
func increaseCounter(cnt *int, c rune) error {
	n, err := strconv.Atoi(string(c))
	if err != nil {
		return err
	}

	*cnt = *cnt*10 + n

	return nil
}
