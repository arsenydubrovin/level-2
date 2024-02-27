package pattern

/*
Паттерн «стратегия» выделяет алгоритмы в отдельные объекты, что позволяет взаимозаменять их во время исполнения.

Плюсы:
- Замена алгоритмов на лету.
- Изолирует код и данные алгоритмов.
- Уход от наследования к делегированию.
Минусы:
- Усложняет программу за счёт дополнительных объектов.
*/

import (
	"fmt"
	"sort"
)

// Интерфейс стратегии сортировки
type sortStrategy interface {
	sort(numbers []int) []int
}

// Сортировка по возрастанию
type ascendingSort struct{}

func (as *ascendingSort) sort(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

// Сортировка по убыванию
type descendingSort struct{}

func (ds *descendingSort) sort(numbers []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	return numbers
}

// Исполнитель стратегии
type sorter struct {
	strategy sortStrategy
}

func (s *sorter) setStrategy(strategy sortStrategy) {
	s.strategy = strategy
}

func (s *sorter) executeStrategy(numbers []int) []int {
	if s.strategy == nil {
		return numbers
	}
	return s.strategy.sort(numbers)
}

func main() {
	numbers := []int{34, 3, -21, 93, 0, 171}

	sorter := &sorter{}

	ascendingSort := &ascendingSort{}
	sorter.setStrategy(ascendingSort)
	fmt.Println("По возврастанию:", sorter.executeStrategy(numbers))

	descendingSort := &descendingSort{}
	sorter.setStrategy(descendingSort)
	fmt.Println("По убыванию:", sorter.executeStrategy(numbers))
}
