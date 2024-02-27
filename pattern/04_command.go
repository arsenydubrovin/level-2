package pattern

/*
Паттерн «комманда» позволяет оборачивать запросы или простые операции в отдельные объекты.

Плюсы:
- Убирает прямую зависимость между операцией и исполнителем.
- Позволяет реализовать историю, отмену, повтор операций.
- Позволяет компоновать сложные команды из простых.
Минусы:
- Усложняет код за счёт введения множества классов.
*/

import "fmt"

// Интерфейс команды
type editorCommand interface {
	execute(editor *textEditor)
}

type textEditor struct {
	text string
}

func newTextEditor() *textEditor {
	return &textEditor{
		text: "",
	}
}

// Конкретная команда
type insertTextCommand struct {
	text string
}

func newInsertTextCommand(text string) *insertTextCommand {
	return &insertTextCommand{text: text}
}

func (c *insertTextCommand) execute(editor *textEditor) {
	editor.text = editor.text + c.text
}

func main() {
	editor := newTextEditor()

	insertCommand := newInsertTextCommand("Привет!")

	insertCommand.execute(editor)

	fmt.Printf("Текст в редакторе:\n%s\n", editor.text)
}
