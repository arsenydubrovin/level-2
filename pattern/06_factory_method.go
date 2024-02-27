package pattern

/*
Паттерн «фабричный метод» позволяет создавать различные объекты, не привязываясь к их классам.

Плюсы:
- Код создания объектов выносится в одно место.
- Упрощает добавление новых объектов.
Минусы
- Может привести к дублированию классов.
*/

import (
	"fmt"
)

type databaseType int

const (
	postgres databaseType = iota
	mongo
)

// Интерфейс базы данных
type database interface {
	connect() error
	query(string)
}

// Конкретная база данных
type postgresDatabase struct{}

func (db *postgresDatabase) connect() error {
	fmt.Println("Соединение с PostgreSQL...")
	return nil
}

func (db *postgresDatabase) query(stmt string) {
	fmt.Printf("Выполение запроса к PostgreSQL: %s\n", stmt)
}

// Конкретная база данных
type mongoDatabase struct{}

func (db *mongoDatabase) connect() error {
	fmt.Println("Соединение с MongoDB...")
	return nil
}

func (db *mongoDatabase) query(stmt string) {
	fmt.Printf("Выполение запроса к MongoDB: %s\n", stmt)
}

// Простая фабрика
type databaseFactory struct{}

func (df *databaseFactory) createDatabase(dbType databaseType) database {
	switch dbType {
	case postgres:
		return &postgresDatabase{}
	case mongo:
		return &mongoDatabase{}
	default:
		return nil
	}
}

func main() {
	factory := &databaseFactory{}

	postgresDB := factory.createDatabase(postgres)
	postgresDB.connect()
	postgresDB.query("SELECT * FROM users")

	fmt.Println()

	mongoDB := factory.createDatabase(mongo)
	mongoDB.connect()
	mongoDB.query("SELECT * FROM events")
}
