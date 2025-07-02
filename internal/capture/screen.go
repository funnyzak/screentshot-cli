package capture

import (
	"fmt"
	"image"
	"runtime"

	"github.com/funnyzak/screenshot-cli/internal/config"
	"github.com/kbinani/screenshot"
)

// CaptureScreen captures a screenshot based on the configuration
func CaptureScreen(config *config.Config) (image.Image, error) {
	if config.Region != nil {
		return captureRegion(config.Region)
	}
	return captureFullScreen(config.Display)
}

// captureFullScreen captures the entire screen using kbinani/screenshot
func captureFullScreen(displayIndex int) (image.Image, error) {
	// Get the number of active displays
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		return nil, fmt.Errorf("no active displays found")
	}

	// Validate display index
	if displayIndex < 0 || displayIndex >= n {
		return nil, fmt.Errorf("invalid display index %d, available displays: %d", displayIndex, n)
	}

	// Capture the specified display
	img, err := screenshot.CaptureDisplay(displayIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to capture display %d: %w", displayIndex, err)
	}

	return img, nil
}

// captureRegion captures a specific region of the screen using kbinani/screenshot
func captureRegion(region *config.Region) (image.Image, error) {
	// Create a rectangle from the region
	rect := image.Rect(region.X, region.Y, region.X+region.Width, region.Y+region.Height)

	// Capture the specified region
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		return nil, fmt.Errorf("failed to capture region: %w", err)
	}

	return img, nil
}

// GetDisplayInfo returns information about available displays
func GetDisplayInfo() ([]image.Rectangle, error) {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		return nil, fmt.Errorf("no active displays found")
	}

	displays := make([]image.Rectangle, n)
	for i := 0; i < n; i++ {
		displays[i] = screenshot.GetDisplayBounds(i)
	}

	return displays, nil
}

// GetDisplayCount returns the number of active displays
func GetDisplayCount() int {
	return screenshot.NumActiveDisplays()
}

// CaptureDisplay captures a specific display by index
func CaptureDisplay(displayIndex int) (image.Image, error) {
	n := screenshot.NumActiveDisplays()
	if displayIndex < 0 || displayIndex >= n {
		return nil, fmt.Errorf("invalid display index %d, available displays: %d", displayIndex, n)
	}

	img, err := screenshot.CaptureDisplay(displayIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to capture display %d: %w", displayIndex, err)
	}

	return img, nil
}

// ValidateRegion checks if a region is valid for the current display setup
func ValidateRegion(region *config.Region) error {
	if region.Width <= 0 || region.Height <= 0 {
		return fmt.Errorf("invalid region dimensions: %dx%d", region.Width, region.Height)
	}

	if region.X < 0 || region.Y < 0 {
		return fmt.Errorf("invalid region position: (%d,%d)", region.X, region.Y)
	}

	// Check if region is within any display bounds
	displays, err := GetDisplayInfo()
	if err != nil {
		// If we can't get display info, just validate basic constraints
		return nil
	}

	regionRect := image.Rect(region.X, region.Y, region.X+region.Width, region.Y+region.Height)

	// Check if the region intersects with any display
	for _, display := range displays {
		if regionRect.Overlaps(display) {
			return nil // Region is valid
		}
	}

	return fmt.Errorf("region (%d,%d,%d,%d) is outside all display bounds",
		region.X, region.Y, region.Width, region.Height)
}

// GetPlatformInfo returns information about the current platform
func GetPlatformInfo() string {
	return fmt.Sprintf("OS: %s, Architecture: %s", runtime.GOOS, runtime.GOARCH)
}

// IsPlatformSupported checks if the current platform is supported
func IsPlatformSupported() bool {
	// kbinani/screenshot supports: windows, darwin, linux, freebsd, openbsd, netbsd
	supportedOS := map[string]bool{
		"windows": true,
		"darwin":  true,
		"linux":   true,
		"freebsd": true,
		"openbsd": true,
		"netbsd":  true,
	}

	return supportedOS[runtime.GOOS]
}
