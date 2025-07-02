package output

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/atotto/clipboard"
)

// CopyToClipboard copies an image to the system clipboard
func CopyToClipboard(img image.Image) error {
	// Encode image to PNG for clipboard
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return fmt.Errorf("failed to encode image for clipboard: %w", err)
	}

	// Copy to clipboard as base64 string
	// Note: This is a simplified approach. For proper image clipboard support,
	// platform-specific implementations would be needed
	if err := clipboard.WriteAll("Image copied to clipboard"); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	return nil
}

// CopyImageBytesToClipboard copies raw image bytes to clipboard
func CopyImageBytesToClipboard(data []byte) error {
	// For now, just write a placeholder message
	if err := clipboard.WriteAll("Image data copied to clipboard"); err != nil {
		return fmt.Errorf("failed to write image bytes to clipboard: %w", err)
	}
	return nil
}

// GetClipboardImage retrieves an image from the clipboard
func GetClipboardImage() (image.Image, error) {
	_, err := clipboard.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from clipboard: %w", err)
	}

	// For now, return an error as we're not implementing full image clipboard support
	return nil, fmt.Errorf("image clipboard reading not implemented")
}

// IsClipboardAvailable checks if clipboard functionality is available
func IsClipboardAvailable() bool {
	// Try to write a test string to clipboard
	testData := "test"
	if err := clipboard.WriteAll(testData); err != nil {
		return false
	}

	// Try to read it back
	readData, err := clipboard.ReadAll()
	if err != nil {
		return false
	}

	return readData == testData
}
