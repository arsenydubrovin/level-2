package pattern

/*
Паттерн строитель» позволяет создавать объекты пошагово. Может производить различные объекты, используя один и тот же процесс строительства.

Плюсы:
- Позволяет создавать объекты пошагово.
- Позволяет может создавать разные объекты.
- Изолирует процесс сборки.
Минусы:
- Усложняет код из-за введения дополнительных классов.
*/

import "fmt"

// Конфиг для HTTP сервера
type httpConfig struct {
	port           int
	address        string
	maxConnections int
	timeout        int
}

// Интефейс строителя
type configBuilder interface {
	setPort(port int) configBuilder
	setAddress(address string) configBuilder
	setMaxConnections(maxConnections int) configBuilder
	setTimeout(timeout int) configBuilder
	build() httpConfig
}

// Конкретный строитель
type httpConfigBuilder struct {
	cfg httpConfig
}

// Конструктор конкретного строителя
func newConfigBuilder() configBuilder {
	return &httpConfigBuilder{}
}

func (b *httpConfigBuilder) setPort(port int) configBuilder {
	b.cfg.port = port
	return b
}

func (b *httpConfigBuilder) setAddress(address string) configBuilder {
	b.cfg.address = address
	return b
}

func (b *httpConfigBuilder) setMaxConnections(maxConnections int) configBuilder {
	b.cfg.maxConnections = maxConnections
	return b
}

func (b *httpConfigBuilder) setTimeout(timeout int) configBuilder {
	b.cfg.timeout = timeout
	return b
}

func (b *httpConfigBuilder) build() httpConfig {
	return b.cfg
}

func main() {
	builder := newConfigBuilder()

	cfg := builder.
		setPort(8080).
		setAddress("localhost").
		setMaxConnections(100).
		setTimeout(30).
		build()

	fmt.Printf("Конфигурация сервера:\n%+v\n", cfg)
}
