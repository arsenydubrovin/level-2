package pattern

/*
Паттер «посетитель» позволяет добавлять объектам новые операции, не изменяя сами объекты.

Плюсы:
- Упрощает добавление новых операций к сложным объектам.
Минусы:
- Если структура объектов часто меняется, паттерн будет мешать.
*/

import "fmt"

// Интерфейс элемента корзины
type element interface {
	accept(v visitor)
}

// Конкретный продукт в корзине
type product struct {
	name  string
	price float64
}

func (p *product) accept(v visitor) {
	v.visitProduct(p)
}

// Корзина с продуктами
type cart struct {
	items []product
}

func (c *cart) accept(v visitor) {
	for _, item := range c.items {
		item.accept(v)
	}
}

// Интерфейс посетителя
type visitor interface {
	visitProduct(*product)
	visitCart(*cart)
}

// Конкретный посетитель для вычисления стоимости всей корзины
type totalPriceCounter struct {
	totalPrice float64
}

func (tc *totalPriceCounter) visitProduct(p *product) {
	fmt.Printf("Посетил продукт %s\n", p.name)
	tc.totalPrice += p.price
}

func (tc *totalPriceCounter) visitCart(c *cart) {
	fmt.Printf("Вычисление...\n")
}

func main() {
	cart := cart{
		items: []product{
			{name: "А", price: 23.49},
			{name: "B", price: 19.89},
			{name: "C", price: 14.98},
		},
	}

	v := &totalPriceCounter{}
	cart.accept(v)

	fmt.Printf("Стоимость всей корзины: %.2f\n", v.totalPrice)
}
