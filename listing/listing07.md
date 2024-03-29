Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:

Программа выведет значения, переданные в `asChan()` в произвольном порядке, а потом бесконечно будет выводить нули.

`asChan()` принимает вариативное число аргументов и вызывает горутину, которая пишет полученные числа в канал и закрывает его.

`merge()` принимает на вход два канала для чтения, в бесконечном цикле читает из каналов и объединяет значения в канал `c`.

`main()` читает и выводит на экран значения из `c`.

После того, как в каналы `a` и `b` запишутся все значения, эти каналы закрываются. Но бесконечный цикл в `merge()` продолжает читать значения из закрытых каналов, в они возвращают нулевое значение типа `int`.

Чтобы остановить бесконечную запись в канал `с`, надо закрывать его, когда каналы `a` и `b` закрыты.

Для этого в операторе `select` закрытый канал нужно приравнивать к `nil`, чтобы остановить чтение и него. А когда оба канала станут равны `nil`, закрывать канал `c` и завершать горутину:

```go
defer close(c)
for {
	select {
	case v, ok := <-a:
		if ok {
			c <- v
		} else {
			a = nil
		}
	case v, ok := <-b:
		if ok {
			c <- v
		} else {
			b = nil
		}
	}
	if a == nil && b == nil {
		return
	}
}
	```
