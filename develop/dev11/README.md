# Сервис для работы с календарём

## Запуск

Запустите приложение:

```bash
make run
```

→ [localhost:1337](http://localhost:1337)

## Локальная разработка

Установите необходимые инструменты: `air`, `gofumpt`, `golangci-lint`, `pre-commit`:

```bash
make init
```

Запустите приложение в live-режиме:

```bash
make run-air
```

Дополнительно:

|   |   |
|---|---|
| `make deps` | Обновить зависимости |
| `make lint` | Запустить линтеры для всего проекта |
| `make` | Показать `help` |

## Примеры запросов

Создать события:

```bash
curl -i -X POST -d "starts_at=2024-02-27T08:00:00Z&ends_at=2024-02-27T08:30:00Z&title=Sample Event&desc=This is a sample event" http://localhost:1337/create_event
```

Обновить событие:

```bash
curl -i -X POST -d "id=1&ends_at=2024-02-27T09:10:00Z" http://localhost:1337/update_event
```

Получить все события на день:

```bash
curl -i "http://localhost:1337/events_for_day?date=2024-02-27"
```

_Выбираются та неделя или месяц, который соответсвуют переданной дате (date)._

На неделю:

```bash
curl -i "http://localhost:1337/events_for_week?date=2024-02-27"
```

На месяц:

```bash
curl -i "http://localhost:1337/events_for_month?date=2024-02-27"
```

Удалить событие:

```bash
curl -i -X POST -d "id=1" http://localhost:1337/delete_event
```