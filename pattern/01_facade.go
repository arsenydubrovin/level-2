package pattern

import "fmt"

/*
Паттерн «фасад» предоставляет клиенту простой интерфейс для сложной системы объектов, библиотеки или фреймворка.

Плюсы:
- Изолирует клиентов от компонентов сложной подсистемы.
Минусы:
- Фасад рискует стать «божественным объектом», привязанным ко всем сущностям программы.
*/

// Сложный интерфейс взаимодействия с базой данных
type Database struct{}

func (d *Database) connect() {
	fmt.Println("Установка соединения...")
}

func (d *Database) initialize() {
	fmt.Println("Инициализация...")
}

func (d *Database) insertUser() int {
	id := 5
	fmt.Printf("Создан пользователь %d\n", id)
	return id
}

func (d *Database) giveAdminPermissions(id int) {
	fmt.Printf("Пользователю с id=%d даны права администратора...\n", id)
}

// Фасад, который предоставляет простой интерфейс
type dbFacade struct {
	db Database
}

func NewDBFacade(db Database) dbFacade {
	return dbFacade{
		db: db,
	}
}

func (df *dbFacade) establishConnection() {
	df.db.connect()
	df.db.initialize()
}

func (df *dbFacade) createAdmin() {
	id := df.db.insertUser()
	df.db.giveAdminPermissions(id)
}

func main() {
	db := Database{}
	facade := NewDBFacade(db)

	facade.establishConnection()
	facade.createAdmin()
}
