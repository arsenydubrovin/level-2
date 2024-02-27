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

Программа напечатает `error`, потому что проверка `err != nil` будет равна `true`.

Тип `customError` соответсвует встроенному интерфейсу `error`. `test()` возвращает `*customError` равный `nil`.

В `main()` объявляется переменная `err`, тип которой — интерфейс `error`. В эту переменную записывается результат `test()`.

Условие `err != nil` проверяет, имееет ли переменная тип нулевой интерфейс (nil interface). Но переменная `err` не равна `nil`, потому что она содержит указатель на тип `*customError` (см. [listing03.md](listing03.md)).

Чтобы проверка работала корректно, можно использовать type assertion: `err.(*customError)` и утвердить тип переменной `err`. Другой способ — функция `Is()` из пакета `errors`: `errors.Is(err, nil)`.