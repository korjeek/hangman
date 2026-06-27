# Hangman Game

🌐 [Read this in Russian / Нажмите здесь для версии на русском языке](README.ru.md)

---

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
