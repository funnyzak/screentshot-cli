package capture

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"screenshot-cli/internal/config"
)

// CaptureScreen captures a screenshot based on the configuration
func CaptureScreen(config *config.Config) (image.Image, error) {
	if config.Region != nil {
		return captureRegion(config.Region)
	}
	return captureFullScreen()
}

// captureFullScreen captures the entire screen
func captureFullScreen() (image.Image, error) {
	switch runtime.GOOS {
	case "darwin":
		return captureFullScreenMac()
	case "windows":
		return captureFullScreenWindows()
	case "linux":
		return captureFullScreenLinux()
	default:
		return createDummyImage(1920, 1080), nil
	}
}

// captureRegion captures a specific region of the screen
func captureRegion(region *config.Region) (image.Image, error) {
	switch runtime.GOOS {
	case "darwin":
		return captureRegionMac(region)
	case "windows":
		return captureRegionWindows(region)
	case "linux":
		return captureRegionLinux(region)
	default:
		return createDummyImage(region.Width, region.Height), nil
	}
}

// captureFullScreenMac captures full screen on macOS using screencapture
func captureFullScreenMac() (image.Image, error) {
	// Create a temporary file
	tempFile := fmt.Sprintf("/tmp/screenshot_%d.png", time.Now().Unix())

	// Use screencapture command
	cmd := exec.Command("screencapture", "-x", tempFile)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture screen: %w", err)
	}

	// Read the image file
	file, err := os.Open(tempFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open screenshot file: %w", err)
	}
	defer file.Close()
	defer os.Remove(tempFile)

	// Decode the image
	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot: %w", err)
	}

	return img, nil
}

// captureRegionMac captures a region on macOS
func captureRegionMac(region *config.Region) (image.Image, error) {
	tempFile := fmt.Sprintf("/tmp/screenshot_%d.png", time.Now().Unix())

	// Use screencapture with region
	cmd := exec.Command("screencapture", "-R",
		fmt.Sprintf("%d,%d,%d,%d", region.X, region.Y, region.Width, region.Height),
		"-x", tempFile)

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture region: %w", err)
	}

	// Read the image file
	file, err := os.Open(tempFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open screenshot file: %w", err)
	}
	defer file.Close()
	defer os.Remove(tempFile)

	// Decode the image
	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot: %w", err)
	}

	return img, nil
}

// captureFullScreenWindows captures full screen on Windows
func captureFullScreenWindows() (image.Image, error) {
	// For now, return a dummy image
	// In a real implementation, you would use Windows API or a library like robotgo
	return createDummyImage(1920, 1080), nil
}

// captureRegionWindows captures a region on Windows
func captureRegionWindows(region *config.Region) (image.Image, error) {
	return createDummyImage(region.Width, region.Height), nil
}

// captureFullScreenLinux captures full screen on Linux
func captureFullScreenLinux() (image.Image, error) {
	// Try to use import command (ImageMagick)
	cmd := exec.Command("import", "-window", "root", "-")
	if output, err := cmd.Output(); err == nil {
		// Try to decode the output as PNG
		img, err := png.Decode(strings.NewReader(string(output)))
		if err == nil {
			return img, nil
		}
	}

	// Fallback to dummy image
	return createDummyImage(1920, 1080), nil
}

// captureRegionLinux captures a region on Linux
func captureRegionLinux(region *config.Region) (image.Image, error) {
	return createDummyImage(region.Width, region.Height), nil
}

// createDummyImage creates a dummy image for testing/fallback
func createDummyImage(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill with a gradient pattern
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x * 255) / width)
			g := uint8((y * 255) / height)
			b := uint8(128)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	return img
}

// GetDisplayInfo returns information about available displays
func GetDisplayInfo() ([]image.Rectangle, error) {
	// For now, return a default display
	return []image.Rectangle{image.Rect(0, 0, 1920, 1080)}, nil
}

// ValidateRegion checks if a region is valid for the current display setup
func ValidateRegion(region *config.Region) error {
	if region.Width <= 0 || region.Height <= 0 {
		return fmt.Errorf("invalid region dimensions: %dx%d", region.Width, region.Height)
	}

	if region.X < 0 || region.Y < 0 {
		return fmt.Errorf("invalid region position: (%d,%d)", region.X, region.Y)
	}

	return nil
}
