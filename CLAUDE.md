# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go microservice for image comparison using a Siamese neural network (ONNX model). The service receives two images via gRPC, processes them through the neural network, and returns a similarity score. Module name is `todo-bichig-comparison` using Go 1.24.

## Project Structure

```
todo-bichig-comparison/
├── cmd/
│   └── app/
│       └── main.go                 # Main application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── server/
│   │   └── grpc.go                # gRPC server implementation
│   └── service/
│       ├── comparison.go          # Image comparison service logic
│       └── image_processor.go     # Image preprocessing (48x48 grayscale)
├── api/
│   └── proto/
│       ├── image_comparison.proto # gRPC API contract
│       └── generated/             # Auto-generated protobuf code
├── models/
│   └── siamese_network.onnx      # Siamese neural network model (163KB, MIT license)
└── go.mod                        # Go module definition
```

## Common Commands

### Generate protobuf code
```bash
protoc --go_out=api/proto/generated/ --go-grpc_out=api/proto/generated/ api/proto/image_comparison.proto
```

### Build and Run
```bash
go run cmd/app/main.go
```

### Build executable
```bash
go build -o todo-bichig-comparison cmd/app/main.go
```

### Module management
```bash
go mod tidy       # Clean up dependencies
go mod download   # Download dependencies
```

### Testing
```bash
go test ./...     # Run all tests
go test -v ./...  # Run tests with verbose output
```

### Code quality
```bash
go fmt ./...      # Format code
go vet ./...      # Run Go vet for potential issues
```

## Architecture Notes

- **gRPC Service**: Accepts two images as bytes and returns similarity score (0.0-1.0)
- **ONNX Integration**: Uses `github.com/microsoft/onnxruntime-go` for model inference
- **Image Processing**: Converts images to 48x48 grayscale format using `github.com/disintegration/imaging`
- **Flow**: API Gateway (HTTP) → gRPC → Image Processing → ONNX Model → Response