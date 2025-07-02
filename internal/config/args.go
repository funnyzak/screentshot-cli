package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Config holds all configuration for the screenshot tool
type Config struct {
	// Screenshot settings
	Region     *Region
	OutputPath string
	Display    int // Display index to capture

	// Output control
	Format    string
	Quality   int
	Clipboard bool
	Template  string

	// Batch processing
	Count    int
	Interval int
	Prefix   string
	Dir      string

	// Internal
	Counter int
}

// Region represents a screenshot region
type Region struct {
	X, Y, Width, Height int
}

// Exit codes
const (
	ExitSuccess        = 0
	ExitArgsError      = 1
	ExitFileError      = 2
	ExitCaptureError   = 3
	ExitFormatError    = 4
	ExitClipboardError = 5
)

// ParseArgs parses command line arguments and returns a Config
func ParseArgs(cmd *cobra.Command, args []string) (*Config, error) {
	config := &Config{}

	// Parse region
	if regionStr, _ := cmd.Flags().GetString("region"); regionStr != "" {
		region, err := parseRegion(regionStr)
		if err != nil {
			return nil, fmt.Errorf("invalid region format: %w", err)
		}
		config.Region = region
	}

	// Parse display
	display, _ := cmd.Flags().GetInt("display")
	if display < 0 {
		return nil, fmt.Errorf("display index must be non-negative")
	}
	config.Display = display

	// Parse output path
	outputPath, _ := cmd.Flags().GetString("output")
	if len(args) > 0 {
		outputPath = args[0]
	}
	config.OutputPath = outputPath

	// Parse format
	format, _ := cmd.Flags().GetString("format")
	if !isValidFormat(format) {
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	config.Format = format

	// Parse quality
	quality, _ := cmd.Flags().GetInt("quality")
	if quality < 1 || quality > 100 {
		return nil, fmt.Errorf("quality must be between 1 and 100")
	}
	config.Quality = quality

	// Parse clipboard
	clipboard, _ := cmd.Flags().GetBool("clipboard")
	config.Clipboard = clipboard

	// Parse template
	template, _ := cmd.Flags().GetString("template")
	config.Template = template

	// Parse batch settings
	count, _ := cmd.Flags().GetInt("count")
	if count < 1 {
		return nil, fmt.Errorf("count must be at least 1")
	}
	config.Count = count

	interval, _ := cmd.Flags().GetInt("interval")
	if interval < 1 {
		return nil, fmt.Errorf("interval must be at least 1 second")
	}
	config.Interval = interval

	prefix, _ := cmd.Flags().GetString("prefix")
	config.Prefix = prefix

	dir, _ := cmd.Flags().GetString("directory")
	config.Dir = dir

	// Process template if provided
	if config.Template != "" {
		config.OutputPath = config.Template
	}

	return config, nil
}

// NewTemplateProcessor creates a new template processor for this config
func (c *Config) NewTemplateProcessor() *TemplateProcessor {
	return NewTemplateProcessor()
}

// parseRegion parses a region string in format "x,y,width,height"
func parseRegion(regionStr string) (*Region, error) {
	parts := strings.Split(regionStr, ",")
	if len(parts) != 4 {
		return nil, fmt.Errorf("region must be in format 'x,y,width,height'")
	}

	coords := make([]int, 4)
	for i, part := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid coordinate: %s", part)
		}
		coords[i] = val
	}

	return &Region{
		X:      coords[0],
		Y:      coords[1],
		Width:  coords[2],
		Height: coords[3],
	}, nil
}

// isValidFormat checks if the format is supported
func isValidFormat(format string) bool {
	validFormats := map[string]bool{
		"png":  true,
		"jpg":  true,
		"jpeg": true,
		"bmp":  true,
		"gif":  true,
	}
	return validFormats[strings.ToLower(format)]
}
