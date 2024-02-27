package pattern

/*
Паттерн «цепочка вызовов» позволяет построить цепь из обработчиков запроса. Каждый обработчик сам решает, может ли он обработать запрос.

Плюсы:
- Уменьшает зависимость между клиентом и обработчиком.
- Реализует принцип single responsibility.
*/

import "fmt"

// Запрос определённого уровня
type request struct {
	level int
}

type handler interface {
	handle(*request)
	setNext(handler)
}

// Хендлер обрабатывает запросы уровня level
type levelHandler struct {
	level       int
	next handler
}

func newLevelHandler(level int) *levelHandler {
	return &levelHandler{level: level}
}

// Логика обработчика
func (h *levelHandler) handle(r *request) {
	fmt.Printf("Хендлер с уровнем %d получил запрос\n", h.level)

	if r.level <= h.level {
		fmt.Printf("Запрос обработан хендлером с уровнем %d\n", h.level)
		return
	}

	if h.next != nil {
		fmt.Printf("Хендлер с уровнем %d не может обработать этот запрос. Запрос передан следующему хендлеру\n", h.level)
		h.next.handle(r)
		return
	}

	fmt.Println("Ни один хендлер не может обработать запрос")
}

func (h *levelHandler) setNext(next handler) {
	h.next = next
}

func main() {
	h1 := newLevelHandler(1)
	h2 := newLevelHandler(3)
	h3 := newLevelHandler(5)

	h1.setNext(h2)
	h2.setNext(h3)

	r1 := &request{level: 3}
	r2 := &request{level: 6}

	fmt.Printf("Обработка запроса с с уровнем %d:\n", r1.level)
	h1.handle(r1)

	fmt.Printf("\nОбработка запроса с уровнем %d:\n", r2.level)
	h1.handle(r2)
}
