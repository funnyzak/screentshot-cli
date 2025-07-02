package output

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"runtime"

	"github.com/atotto/clipboard"
)

// CopyToClipboard copies an image to the system clipboard
func CopyToClipboard(img image.Image) error {
	// Encode image to PNG for clipboard
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return fmt.Errorf("failed to encode image for clipboard: %w", err)
	}

	imgData := buf.Bytes()

	// Try platform-specific image clipboard support first
	if err := copyImageToClipboardPlatform(imgData); err == nil {
		return nil
	}

	// Fallback: try to copy as base64 data URL
	if err := copyImageAsDataURL(imgData); err == nil {
		return nil
	}

	// Last resort: save to temp file and copy file path
	return copyImageViaTempFile(imgData)
}

// copyImageToClipboardPlatform attempts platform-specific image clipboard copying
func copyImageToClipboardPlatform(imgData []byte) error {
	switch runtime.GOOS {
	case "darwin":
		return copyImageToClipboardMacOS(imgData)
	case "windows":
		return copyImageToClipboardWindows(imgData)
	case "linux":
		return copyImageToClipboardLinux(imgData)
	default:
		return fmt.Errorf("platform %s not supported for image clipboard", runtime.GOOS)
	}
}

// copyImageToClipboardMacOS copies image to clipboard on macOS
func copyImageToClipboardMacOS(imgData []byte) error {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "screenshot-*.png")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Write image data to temp file
	if _, err := tmpFile.Write(imgData); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}
	tmpFile.Close()

	// Use osascript to copy the file to clipboard
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`
		set the clipboard to (read POSIX file "%s" as JPEG picture)
	`, tmpFile.Name()))

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy image to clipboard via osascript: %w", err)
	}

	return nil
}

// copyImageToClipboardWindows copies image to clipboard on Windows
func copyImageToClipboardWindows(imgData []byte) error {
	// For Windows, we'll use PowerShell to copy image to clipboard
	// Convert image data to base64
	base64Data := base64.StdEncoding.EncodeToString(imgData)

	psScript := fmt.Sprintf(`
		Add-Type -AssemblyName System.Windows.Forms
		$image = [System.Drawing.Image]::FromStream([System.IO.MemoryStream][System.Convert]::FromBase64String('%s'))
		$bitmap = New-Object System.Drawing.Bitmap($image)
		[System.Windows.Forms.Clipboard]::SetImage($bitmap)
		$image.Dispose()
		$bitmap.Dispose()
	`, base64Data)

	cmd := exec.Command("powershell", "-Command", psScript)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy image to clipboard via PowerShell: %w", err)
	}

	return nil
}

// copyImageToClipboardLinux copies image to clipboard on Linux
func copyImageToClipboardLinux(imgData []byte) error {
	// For Linux, try xclip with image format
	cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "image/png", "-i", "-")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start xclip: %w", err)
	}

	if _, err := stdin.Write(imgData); err != nil {
		return fmt.Errorf("failed to write image data to xclip: %w", err)
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("xclip failed: %w", err)
	}

	return nil
}

// copyImageAsDataURL copies image as a data URL to clipboard
func copyImageAsDataURL(imgData []byte) error {
	base64Data := base64.StdEncoding.EncodeToString(imgData)
	dataURL := fmt.Sprintf("data:image/png;base64,%s", base64Data)

	if err := clipboard.WriteAll(dataURL); err != nil {
		return fmt.Errorf("failed to write data URL to clipboard: %w", err)
	}

	return nil
}

// copyImageViaTempFile saves image to temp file and copies file path
func copyImageViaTempFile(imgData []byte) error {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "screenshot-*.png")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Write image data to temp file
	if _, err := tmpFile.Write(imgData); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Copy file path to clipboard as fallback
	message := fmt.Sprintf("Screenshot saved to: %s", tmpFile.Name())
	if err := clipboard.WriteAll(message); err != nil {
		return fmt.Errorf("failed to write file path to clipboard: %w", err)
	}

	return nil
}

// CopyImageBytesToClipboard copies raw image bytes to clipboard
func CopyImageBytesToClipboard(data []byte) error {
	// Try to decode as image first
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		// If it's not a valid image, try copying as base64
		return copyImageAsDataURL(data)
	}

	// If it's a valid image, use the main function
	return CopyToClipboard(img)
}

// GetClipboardImage retrieves an image from the clipboard
func GetClipboardImage() (image.Image, error) {
	// Try to read from clipboard
	data, err := clipboard.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from clipboard: %w", err)
	}

	// Check if it's a data URL
	if len(data) > 10 && data[:10] == "data:image" {
		// Extract base64 data from data URL
		// Format: data:image/png;base64,<base64data>
		parts := bytes.SplitN([]byte(data), []byte(","), 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid data URL format")
		}

		imgData, err := base64.StdEncoding.DecodeString(string(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 data: %w", err)
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			return nil, fmt.Errorf("failed to decode image data: %w", err)
		}

		return img, nil
	}

	return nil, fmt.Errorf("no image data found in clipboard")
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

// IsImageClipboardSupported checks if image clipboard is supported on current platform
func IsImageClipboardSupported() bool {
	switch runtime.GOOS {
	case "darwin":
		// Check if osascript is available
		_, err := exec.LookPath("osascript")
		return err == nil
	case "windows":
		// Check if PowerShell is available
		_, err := exec.LookPath("powershell")
		return err == nil
	case "linux":
		// Check if xclip is available
		_, err := exec.LookPath("xclip")
		return err == nil
	default:
		return false
	}
}
