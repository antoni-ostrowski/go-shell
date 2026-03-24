# go-shell

A basic shell implementation in Go.

## Features

- Built-in commands
- Fallback to system commands for anything else
- Displays current working directory in prompt

## Usage

```bash
go run .
```

## Structure

- `main.go` - Shell loop and input handling
- `program.go` - Program interface for built-in commands
- `programs/` - Built-in command implementations

