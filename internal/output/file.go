package output

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"screenshot-cli/internal/config"
)

// SaveToFile saves an image to a file with the specified configuration
func SaveToFile(img image.Image, config *config.Config) error {
	// Ensure output directory exists
	if err := ensureDirectory(config.OutputPath); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Encode image to the specified format
	data, err := EncodeImage(img, ImageFormat(config.Format), config.Quality)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	// Write to file
	if err := os.WriteFile(config.OutputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// ensureDirectory ensures the directory for the output file exists
func ensureDirectory(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir == "." || dir == "" {
		return nil // Current directory, no need to create
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	return nil
}

// GetOutputPath generates the output path based on configuration and template
func GetOutputPath(config *config.Config, templateProcessor *config.TemplateProcessor) string {
	if config.Template != "" {
		return templateProcessor.ProcessTemplate(config.Template, config)
	}

	// If no template, use the configured output path
	if config.OutputPath != "" {
		return config.OutputPath
	}

	// Default fallback
	return "screenshot.png"
}

// ValidateOutputPath checks if the output path is valid and writable
func ValidateOutputPath(outputPath string) error {
	dir := filepath.Dir(outputPath)

	// Check if directory exists or can be created
	if dir != "." && dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// Try to create the directory
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("cannot create directory %s: %w", dir, err)
			}
		}
	}

	// Check if we can write to the directory
	if dir != "." && dir != "" {
		testFile := filepath.Join(dir, ".test_write")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			return fmt.Errorf("cannot write to directory %s: %w", dir, err)
		}
		os.Remove(testFile) // Clean up test file
	}

	return nil
}
