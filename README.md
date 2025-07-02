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

> When no template is provided, filenames will be auto-numbered (e.g., screenshot_001.png, screenshot_002.png).

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
| `--directory, -d` | Output directory | . |

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

### Basic Screenshots
```bash
# Quick full screen capture (saves as screenshot.png)
sshot

# Full screen with custom filename
sshot -o my_screenshot.png

# Region capture (x, y, width, height)
sshot -r "100,100,800,600" -o region.png

# High quality JPEG for web sharing
sshot -f jpg -q 95 -o web_screenshot.jpg

# BMP format for legacy systems
sshot -f bmp -o legacy_screenshot.bmp

# GIF format for animated content
sshot -f gif -o animated_screenshot.gif
```

### Clipboard Operations
```bash
# Copy to clipboard only (no file saved)
sshot -c

# Copy to clipboard and save file
sshot -c -o saved_and_copied.png

# Copy region to clipboard
sshot -r "0,0,1920,1080" -c

# Copy high quality JPEG to clipboard
sshot -f jpg -q 90 -c
```

### Template System
```bash
# Basic template with datetime
sshot -t "screenshot_{datetime}.png"

# Template with date and counter
sshot -t "screen_{date}_{counter}.jpg" -f jpg

# Template with random string
sshot -t "capture_{random}.png"

# Template with prefix variable
sshot -t "{prefix}_{timestamp}.png" -p "test"

# Complex template with multiple variables
sshot -t "screenshot_{date}_{time}_{random}_{counter}.png"
```

### Batch Processing
```bash
# Quick batch (10 screenshots, 1 second apart)
sshot -n 10 -i 1

# Batch with custom directory
sshot -n 5 -i 2 -d "./screenshots"

# Batch region screenshots
sshot -r "0,0,1920,1080" -n 5 -i 2 -t "monitor_{counter}.png"

# Batch with template and custom directory
sshot -n 20 -i 1 -d "./captures/{date}" -t "shot_{time}_{counter}.png"

# High frequency batch (every 0.5 seconds)
sshot -n 100 -i 0.5 -t "rapid_{counter}.png"

# Long-term monitoring (every 5 minutes for 2 hours)
sshot -n 24 -i 300 -d "./monitoring" -t "monitor_{time}.png"
```

### Development Workflow
```bash
# Take screenshots during development (every 5 seconds)
sshot -n 50 -i 5 -d "./dev-screenshots" -t "dev_{datetime}_{counter}.png"

# Capture specific region for UI testing
sshot -r "0,0,1200,800" -n 10 -i 3 -t "ui_test_{counter}.png"

# Development with high quality JPEG
sshot -r "0,0,1920,1080" -f jpg -q 95 -n 20 -i 2 -t "dev_{date}_{counter}.jpg"
```

### Documentation & Tutorials
```bash
# Create documentation screenshots
sshot -r "0,0,1200,800" -f jpg -q 85 -t "doc_{date}_{counter}.jpg"

# Tutorial screenshots with region capture
sshot -r "100,100,800,600" -n 5 -i 10 -t "tutorial_step_{counter}.png"

# Documentation with custom prefix
sshot -p "doc" -n 10 -i 5 -t "{prefix}_{datetime}_{counter}.png"
```

### Monitoring & Surveillance
```bash
# Monitor system every minute for an hour
sshot -n 60 -i 60 -d "./monitoring" -t "monitor_{time}.png"

# Continuous monitoring (every 30 seconds)
sshot -n 120 -i 30 -d "./surveillance" -t "surv_{datetime}.png"

# High-resolution monitoring
sshot -r "0,0,1920,1080" -f jpg -q 90 -n 100 -i 60 -t "monitor_{time}.jpg"
```

### Web & Social Media
```bash
# Optimized for web sharing (JPEG, medium quality)
sshot -f jpg -q 80 -o web_ready.jpg

# Social media optimized (square region, high quality)
sshot -r "0,0,1080,1080" -f jpg -q 90 -o social_media.jpg

# Web documentation (PNG for transparency)
sshot -r "0,0,1200,800" -f png -o web_doc.png
```

### Gaming & Entertainment
```bash
# Capture game window (adjust coordinates as needed)
sshot -r "0,0,1920,1080" -f png -o game_screenshot.png

# Gaming highlights (high quality)
sshot -r "0,0,1920,1080" -f jpg -q 95 -o gaming_highlight.jpg

# Game recording simulation (rapid capture)
sshot -r "0,0,1920,1080" -n 300 -i 0.1 -t "game_{counter}.png"
```

### Professional Use Cases
```bash
# Presentation screenshots
sshot -r "0,0,1920,1080" -f jpg -q 90 -t "presentation_{date}_{counter}.jpg"

# Meeting recordings (every 30 seconds)
sshot -n 120 -i 30 -d "./meetings" -t "meeting_{datetime}_{counter}.png"

# Training material creation
sshot -r "0,0,1200,800" -n 20 -i 5 -t "training_{date}_{counter}.png"
```

### Advanced Combinations
```bash
# Clipboard + file save + template
sshot -c -t "backup_{datetime}.png"

# Region + clipboard + high quality
sshot -r "100,100,800,600" -c -f jpg -q 95

# Batch + clipboard + custom directory
sshot -n 5 -i 2 -c -d "./temp" -t "batch_{counter}.png"

# Multi-format batch processing
sshot -n 10 -i 1 -f jpg -q 85 -t "multi_{counter}.jpg"
```

### Troubleshooting & Debug
```bash
# Verbose mode for debugging
sshot -c --verbose

# Test clipboard functionality
sshot -c -r "0,0,100,100"

# Test different formats
sshot -f png -o test.png && sshot -f jpg -o test.jpg && sshot -f bmp -o test.bmp

# Test region capture
sshot -r "0,0,100,100" -o test_region.png
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