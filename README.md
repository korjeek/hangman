# Hangman Game / Игра "Виселица"

[![en](https://shields.io)](#english-version)
[![ru](https://shields.io)](#русская-версия)

---

## English Version

A console-based implementation of the classic **Hangman** game written in Go (Golang). The project is built following clean architecture principles, SOLID design, and features robust Black-Box unit testing with high code coverage.

### Key Features
* **Two Execution Modes**: 
  * **Interactive**: Play directly in the terminal with random words selected from an embedded JSON dictionary based on chosen categories and difficulty levels.
  * **Non-Interactive**: Automated simulation mode that evaluates predefined word and guess inputs (useful for integrations and automated pipelines).
* **SOLID Architecture**: Clean separation between the domain logic, infrastructure data stores, and user interface layer via Go interfaces.
* **Deterministic Pseudo-Random Generation**: Abstracted random selection allows predictable unit testing via mocks without relying on mutable global states.
* **Comprehensive Testing**: 
  * Strict **Black-Box** tests (`package ..._test`).
  * Automated interface mocking utilizing Uber's `mockgen`.
  * Safe multi-threaded test execution (`t.Parallel()`).
  * Detailed code coverage profiling (90%+ domain coverage).

### Prerequisites
* **Go**: Version 1.22 or higher is recommended.
* **Mockgen** (Only required if you plan to regenerate interface mocks):
  ```bash
  go install go.uber.org/mock/mockgen@latest
  ```

### Installation & Execution
1. Clone the repository and navigate to the project root directory.
2. Build and run the game in **Interactive Mode** (default):
   ```bash
   go run cmd/hangman/main.go
   ```
3. Run the game with specific optional flags (e.g., set category or difficulty):
   ```bash
   go run cmd/hangman/main.go --category animals --difficulty hard
   ```
4. Run the game in **Non-Interactive Mode** by passing positional arguments:
   ```bash
   go run cmd/hangman/main.go "secretword" "aeiou"
   ```

### Testing & Code Coverage
To run the test suite and verify the code coverage statistics, execute the following commands in your terminal:
```bash
# Generate interface mocks
go generate ./...

# Run tests and save the coverage profile
go test -coverprofile="coverage.out" ./...

# View detailed code coverage breakdown in your browser
go tool cover -html="coverage.out"
```

---

## Русская версия

Консольная реализация классической игры **"Виселица"**, написанная на языке Go (Golang). Проект разработан в соответствии с принципами чистой архитектуры, SOLID и укомплектован надежными модульными Black-Box тестами с высоким процентом покрытия кода.

### Основные возможности
* **Два режима работы**:
  * **Интерактивный**: Полноценная игра в терминале со случайным выбором слов из встроенного JSON-словаря на основе выбранной категории и уровня сложности.
  * **Неинтерактивный**: Режим автоматической симуляции, который принимает готовую пару слов и последовательность ходов (полезно для интеграций и автоматических проверок).
* **Архитектура SOLID**: Строгое разделение между доменной логикой, инфраструктурными репозиториями данных и слоем пользовательского интерфейса с помощью интерфейсов Go.
* **Детерминированная генерация**: Абстрагирование генератора случайных чисел позволяет полностью контролировать «рандом» в тестах с помощью моков.
* **Глубокое тестирование**:
  * Строгие **Black-Box** тесты (`package ..._test`).
  * Автоматическая генерация моков интерфейсов с помощью утилиты `mockgen` от Uber.
  * Безопасный параллельный запуск тестов (`t.Parallel()`).
  * Детальный расчет покрытия кода (более 90% для доменного слоя).

### Системные требования
* **Go**: Рекомендуется версия 1.22 или выше.
* **Mockgen** (Нужен только в том случае, если вы планируете заново генерировать моки интерфейсов):
  ```bash
  go install go.uber.org/mock/mockgen@latest
  ```

### Сборка и запуск
1. Клонируйте репозиторий и перейдите в корневую папку проекта.
2. Соберите и запустите игру в **Интерактивном режиме** (по умолчанию):
   ```bash
   go run cmd/hangman/main.go
   ```
3. Запустите игру с определенными флагами (например, выберите категорию или сложность):
   ```bash
   go run cmd/hangman/main.go --category animals --difficulty hard
   ```
4. Запустите игру в **Неинтерактивном режиме**, передав позиционные аргументы:
   ```bash
   go run cmd/hangman/main.go "secretword" "aeiou"
   ```

### Тестирование и покрытие кода
Чтобы запустить набор тестов и проверить статистику покрытия кода, выполните следующие команды в терминале:
```bash
# Сгенерировать моки интерфейсов
go generate ./...

# Запустить тесты и сохранить профиль покрытия
go test -coverprofile="coverage.out" ./...

# Открыть детальный интерактивный отчет о покрытии в браузере
go tool cover -html="coverage.out"
```
