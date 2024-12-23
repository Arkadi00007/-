# Сервис подсчёта арифметических выражений|LMS 

## Описание

LMS Calculator Application — это серверное приложение на Go, предназначенное для вычисления математических выражений. Приложение предоставляет консольный интерфейс для ввода выражений и HTTP API для удалённых вычислений.

---
## Как запустить
#### go run cmd/calc_service/main.go
#### Можно поменять порт в main.go(по умолчанию :8080)

## Как использовать

### 1. Консольный режим:
- Пользователь может вводить математические выражения через консоль.
- Поддерживается команда `exit` для выхода из приложения.
- Результат вычисления или ошибка отображаются в консоли.

### 2. HTTP API:
- Путь: `/api/v1/calculate`
- Метод: `POST`
- Формат запроса: JSON.
- Формат ответа: JSON.
- Поддерживает валидацию выражений.

---

## API Документация

### **Запрос**

#### Метод: `POST`
#### URL: `/api/v1/calculate`
#### Заголовки:
- `Content-Type: application/json`

## Тело запроса:
{
"expression": "(здесь можно использовать цифры,скобки,математические операции)"}

#### Пример тела запроса:
```json
{
  "expression": "2+2"
}
```

# Примеры запросов:
1.Запрос
```
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression":"5 + 3 * 2"}'
```
Ответ:
```
{
  "result": 11
}
```

2.Запрос
```
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression":""}'
```
Ответ:
```
{
  "error": "Empty expression"
}
```

3.Запрос
```
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression":"5 + invalid"}'
```
Ответ:
```
{
  "error": "Expression is not valid"
}
```

4.Запрос
```
curl -X GET http://localhost:8080/api/v1/calculate
```
Ответ:
```
{
"error": "Only POST method is allowed"
}
```


