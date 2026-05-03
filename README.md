# 📝 go-tui-todo

A simple, interactive terminal-based todo list built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) — a Go framework for building elegant TUI (Terminal UI) applications.

---

## ✨ Features

- Add new todo items
- Mark todos as complete / incomplete
- Delete todos
- Navigate with keyboard
- Clean and minimal terminal UI

---

## 📦 Tech Stack

| Tool                                                     | Purpose                                    |
| -------------------------------------------------------- | ------------------------------------------ |
| [Go](https://golang.org/)                                | Programming language                       |
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | TUI framework                              |
| [Bubbles](https://github.com/charmbracelet/bubbles)      | Reusable TUI components (text input, list) |
| [Lip Gloss](https://github.com/charmbracelet/lipgloss)   | Terminal styling                           |

---

## 🚀 Getting Started

### Prerequisites

- Go `1.21+` installed → [Download Go](https://golang.org/dl/)

### Installation

```bash
# Clone the repository
git clone https://github.com/Franciss-prog/go-tui-todo.git
cd go-tui-todo

# Install dependencies
go mod tidy

# Run the app
go run .
```

---

## ⌨️ Keybindings

| Key            | Action                       |
| -------------- | ---------------------------- |
| `↑` / `↓`      | Navigate todos               |
| `Enter`        | Toggle complete / incomplete |
| `a`            | Add a new todo               |
| `d`            | Delete selected todo         |
| `q` / `Ctrl+C` | Quit                         |

---

## 📁 Project Structure

```
go-tui-todo/
├── main.go         # Entry point
├── model.go        # Bubble Tea model (state, update, view)
├── todo.go         # Todo item struct and helpers
├── styles.go       # Lip Gloss styles
├── go.mod
└── go.sum
```

---

## 🛠️ How It Works

This app follows the **Elm Architecture** pattern used by Bubble Tea:

- **Model** — holds the application state (list of todos, cursor position, input mode)
- **Update** — handles key events and updates the model
- **View** — renders the current state to the terminal

---

## 📸 Preview

```
┌─────────────────────────────┐
│        📝 My Todos          │
├─────────────────────────────┤
│  ▶ [ ] Buy groceries        │
│    [✓] Read Go docs         │
│    [ ] Build a TUI app      │
├─────────────────────────────┤
│ a: add  d: delete  q: quit  │
└─────────────────────────────┘
```

---

## 📄 License

MIT License — feel free to use and modify.

---
