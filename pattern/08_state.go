package pattern

/*
Паттерн «состояние» позволяет динамически изменять поведение объекта при смене его состояния.

Плюсы:
- Избавляет от сложных ветвлений для управления состояниями.
Минусы:
- Если состояний мало, может сильно усложнить код.
*/

import "fmt"

// Интерфейс состояния посылки
type packageState interface {
	processPackage(pkg *packageStruct)
}

// Посылка хранящее состояние
type packageStruct struct {
	state packageState
}

func (p *packageStruct) setState(state packageState) {
	p.state = state
}

func (p *packageStruct) processPackage() {
	p.state.processPackage(p)
}

type processingState struct{}

func (ps *processingState) processPackage(pkg *packageStruct) {
	fmt.Println("Заказ в обработке.")
}

// outForDeliveryState - конкретное состояние посылки: В пути к получателю.
type outForDeliveryState struct{}

func (ofds *outForDeliveryState) processPackage(pkg *packageStruct) {
	fmt.Println("Заказ в пути.")
}

// deliveredState - конкретное состояние посылки: Доставлено.
type deliveredState struct{}

func (ds *deliveredState) processPackage(pkg *packageStruct) {
	fmt.Println("Заказ доставлен.")
}

func main() {
	packageDelivery := &packageStruct{}

	fmt.Println("Состояния заказа:")

	packageDelivery.setState(&processingState{})
	packageDelivery.processPackage()

	packageDelivery.setState(&outForDeliveryState{})
	packageDelivery.processPackage()

	packageDelivery.setState(&deliveredState{})
	packageDelivery.processPackage()
}
