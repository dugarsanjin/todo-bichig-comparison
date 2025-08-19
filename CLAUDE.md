# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go project named `todo-bichig-comparsion` using Go 1.24. The project appears to be in very early development stages with minimal code structure.

## Project Structure

- `cmd/app/main.go` - Main application entry point (currently empty)
- `go.mod` - Go module definition

## Common Commands

### Build and Run
```bash
go run cmd/app/main.go
```

### Build executable
```bash
go build -o todo-bichig-comparsion cmd/app/main.go
```

### Module management
```bash
go mod tidy    # Clean up dependencies
go mod download # Download dependencies
```

### Testing
```bash
go test ./...  # Run all tests
go test -v ./... # Run tests with verbose output
```

### Code quality
```bash
go fmt ./...   # Format code
go vet ./...   # Run Go vet for potential issues
```

## Architecture Notes

The project follows a standard Go project layout with the main application in `cmd/app/`. Currently, the codebase is minimal with just an empty main function, suggesting this is a new project ready for implementation.