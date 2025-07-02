# Screenshot-CLI

A lightweight, dependency-free Windows command line screenshot tool that supports full screen and region screenshots with flexible output control and batch processing capabilities.

## Features
A lightweight, dependency-free Windows command line screenshot tool that supports full screen and region screenshots with flexible output control and batch processing capabilities.

- **Full Screen Screenshots**: Capture entire screen
- **Region Screenshots**: Capture specific areas with coordinates
- **Multiple Formats**: PNG, JPG, BMP, GIF support
- **Quality Control**: Adjustable JPEG compression quality
- **Clipboard Support**: Copy screenshots directly to clipboard
- **Batch Processing**: Take multiple screenshots with configurable intervals
- **Template System**: Dynamic filename generation with variables
- **Cross-Platform**: Works on Windows, macOS, and Linux

## Installation

### Prerequisites
- Go 1.21 or later

### Build from Source
```bash
git clone git@github.com/funnyzak/screenshot-cli.git
cd screenshot-cli
go build -o sshot cmd/main.go
```

### Install Globally
```bash
go install ./cmd/main.go
```

## Usage

### Basic Screenshots

```bash
# Full screen screenshot (saves as screenshot.png)
sshot

# Full screen screenshot with custom filename
sshot -o desktop.png

# Region screenshot
sshot -r "100,100,800,600" -o region.png

# Specify format and quality
sshot -f jpg -q 80 -o screen.jpg
```

### Output Control

```bash
# Copy to clipboard only
sshot -c

# Use filename template
sshot -t "screenshot_{datetime}.png"

# Multiple template variables
sshot -t "screen_{date}_{counter}.jpg" -f jpg
```

### Batch Processing

```bash
# Take 10 screenshots every 3 seconds
sshot -n 10 -i 3 -p "batch" -d "./screenshots"

# Batch region screenshots
sshot -r "0,0,1920,1080" -n 5 -i 2 -t "monitor_{counter}.png"

# Batch with custom directory structure
sshot -n 20 -i 1 -d "./captures/{date}" -t "shot_{time}_{counter}.png"
```

## Command Line Options

### Global Options
| Option | Description |
|--------|-------------|
| `--help, -h` | Show help information |
| `--version, -v` | Show version information |
| `--verbose` | Enable verbose output |

### Screenshot Options
| Option | Description | Default |
|--------|-------------|---------|
| `--region, -r` | Region screenshot "x,y,width,height" | - |
| `--output, -o` | Output file path | screenshot.png |

### Output Control
| Option | Description | Default |
|--------|-------------|---------|
| `--format, -f` | Output format (png/jpg/bmp/gif) | png |
| `--quality, -q` | JPG compression quality (1-100) | 90 |
| `--clipboard, -c` | Copy to clipboard | false |
| `--template, -t` | Filename template | - |

### Batch Processing
| Option | Description | Default |
|--------|-------------|---------|
| `--count, -n` | Number of screenshots | 1 |
| `--interval, -i` | Screenshot interval (seconds) | 1 |
| `--prefix, -p` | Filename prefix | shot |
| `--dir, -d` | Output directory | . |

## Template Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `{timestamp}` | Unix timestamp | 1640995200 |
| `{datetime}` | Date and time | 20220101_120000 |
| `{date}` | Date only | 20220101 |
| `{time}` | Time only | 120000 |
| `{counter}` | Sequence number | 001, 002, 003 |
| `{random}` | Random string | a1b2c3 |
| `{prefix}` | Filename prefix | shot |

## Examples

### Development Workflow
```bash
# Take screenshots during development
sshot -n 50 -i 5 -d "./dev-screenshots" -t "dev_{datetime}_{counter}.png"
```

### Documentation
```bash
# Create documentation screenshots
sshot -r "0,0,1200,800" -f jpg -q 85 -t "doc_{date}_{counter}.jpg"
```

### Monitoring
```bash
# Monitor system every minute for an hour
sshot -n 60 -i 60 -d "./monitoring" -t "monitor_{time}.png"
```

## Error Handling

The tool provides clear error messages and appropriate exit codes:

- `0`: Success
- `1`: Argument error
- `2`: File system error
- `3`: Screenshot capture error
- `4`: Format conversion error
- `5`: Clipboard error

## Performance

- **Response Time**: < 500ms for full screen, < 300ms for region
- **Memory Usage**: < 50MB
- **File Size**: < 10MB executable
- **CPU Usage**: < 20% during capture

## Development

### Project Structure
```
screenshot-cli/
├── cmd/
│   └── main.go                 # Program entry point
├── internal/
│   ├── capture/
│   │   └── screen.go          # Screenshot core logic
│   ├── output/
│   │   ├── format.go          # Format conversion
│   │   ├── file.go            # File output
│   │   └── clipboard.go       # Clipboard operations
│   ├── batch/
│   │   └── processor.go       # Batch processing logic
│   └── config/
│       ├── args.go            # Command line argument parsing
│       └── template.go        # Filename template processing
├── go.mod
└── README.md
```

### Building
```bash
# Development build
go build -o sshot cmd/main.go

# Release build (Windows)
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o sshot.exe cmd/main.go

# Release build (macOS)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o sshot cmd/main.go

# Release build (Linux)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o sshot cmd/main.go
```

### Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./internal/capture
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## Acknowledgments

- [kbinani/screenshot](https://github.com/kbinani/screenshot) - Cross-platform screenshot library
- [spf13/cobra](https://github.com/spf13/cobra) - Command line framework
- [atotto/clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard library 