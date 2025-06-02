# 🛠️ Go Compiler

A simple compiler written in Go, designed for educational and experimental purposes. It features a hand-written lexer, parser, and evaluator for a basic scripting language.

## ✨ Features to be added.

* Tokenizer (lexer) for source code
* Recursive descent parser
* AST (Abstract Syntax Tree) generation
* REPL for live testing
* Interpreter/evaluator for expressions and statements
* Written 100% in Go with no external dependencies

## 📂 Project Structure

```
/compiler
│
├── lexer/         # Tokenizing logic
├── parser/        # Parsing expressions/statements into AST
├── ast/           # AST node definitions
├── token/         # Token type definitions
├── evaluator/     # Executes parsed AST
├── repl/          # REPL interface
└── main.go        # Entry point
```

## 🚀 Getting Started

### Prerequisites

* Go 1.18 or higher

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-compiler.git
cd go-compiler
```

2. Run the REPL:

```bash
go run main.go
```

You should see:

```
>> let five = 5;
>> five + 10;
15
```

## 💡 Example Code

```js
let name = "Mohit";
let age = 25;
age + 5;
```

## 🧪 Running Tests

Each component (lexer, parser, etc.) has its own `_test.go` files.

To run all tests:

```bash
go test ./...
```

## 📌 TODO

* Add support for functions
* Support conditionals (if/else)
* Handle escape characters in strings
* Add error recovery in parser
* Add support for arrays and hash maps

## 📄 License

MIT License

---

Built as a learning project inspired by the [Monkey Programming Language](https://interpreterbook.com) by Thorsten Ball.
