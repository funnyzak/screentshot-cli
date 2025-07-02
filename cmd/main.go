package main

import (
	"fmt"
	"os"

	"github.com/funnyzak/screenshot-cli/internal/batch"
	"github.com/funnyzak/screenshot-cli/internal/capture"
	"github.com/funnyzak/screenshot-cli/internal/config"
	"github.com/funnyzak/screenshot-cli/internal/output"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	verbose bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "sshot [flags] [output_file]",
		Short: "A lightweight screenshot CLI tool",
		Long: `A lightweight, cross-platform command line screenshot tool that supports full screen and region screenshots with flexible output control and batch processing capabilities.

EXAMPLES:
  # Basic screenshots
  sshot                                    # Full screen screenshot
  sshot -o desktop.png                     # Save with custom filename
  sshot -r "100,100,800,600" -o region.png # Region screenshot
  
  # Output control
  sshot -f jpg -q 80 -o screen.jpg         # JPEG with quality control
  sshot -c                                 # Copy to clipboard only
  sshot -t "screenshot_{datetime}.png"     # Use filename template
  
  # Batch processing
  sshot -n 10 -i 3 -p "batch"              # 10 screenshots every 3 seconds
  sshot -n 5 -i 2 -d "./screenshots"       # Save to custom directory
  
  # Multi-display support
  sshot -d 0 -o primary.png                # Primary display
  sshot -d 1 -o secondary.png              # Secondary display
  
  # Advanced templates
  sshot -t "screen_{date}_{time}_{counter}.png" -n 5
  sshot -d "./captures/{date}" -t "{time}_{random}.png"

TEMPLATE VARIABLES:
  {timestamp}  - Unix timestamp
  {datetime}   - Date and time (YYYYMMDD_HHMMSS)
  {date}       - Date only (YYYYMMDD)
  {time}       - Time only (HHMMSS)
  {counter}    - Sequence number
  {random}     - Random string
  {prefix}     - Filename prefix

SUPPORTED FORMATS:
  png  - Portable Network Graphics (default)
  jpg  - JPEG with quality control
  bmp  - Bitmap
  gif  - Graphics Interchange Format`,
		Version: version,
		RunE:    runScreenshot,
		Example: `  # Quick full screen capture
  sshot

  # Region capture with custom output
  sshot -r "0,0,1920,1080" -o fullscreen.png

  # High quality JPEG for web
  sshot -f jpg -q 95 -o web_screenshot.jpg

  # Copy to clipboard for sharing
  sshot -c

  # Batch capture for monitoring
  sshot -n 60 -i 60 -t "monitor_{time}.png"

  # Development workflow
  sshot -r "0,0,1200,800" -n 10 -i 30 -t "dev_{counter}.png"`,
	}

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "Enable verbose output mode for debugging")

	// Screenshot flags
	rootCmd.Flags().StringP("region", "r", "", "Capture specific region \"x,y,width,height\" (e.g., \"100,100,800,600\")")
	rootCmd.Flags().StringP("output", "o", "screenshot.png", "Output file path (use \"\" for clipboard only)")
	rootCmd.Flags().IntP("display", "d", 0, "Display index to capture (0=primary, 1=secondary, etc.)")

	// Output control flags
	rootCmd.Flags().StringP("format", "f", "png", "Output format: png, jpg, bmp, or gif")
	rootCmd.Flags().IntP("quality", "q", 90, "JPEG compression quality (1-100, higher=better quality)")
	rootCmd.Flags().BoolP("clipboard", "c", false, "Copy screenshot to clipboard")
	rootCmd.Flags().StringP("template", "t", "", "Filename template with variables (e.g., \"{datetime}_{counter}.png\")")

	// Batch processing flags
	rootCmd.Flags().IntP("count", "n", 1, "Number of screenshots to capture (use >1 for batch mode)")
	rootCmd.Flags().IntP("interval", "i", 1, "Interval between screenshots in seconds")
	rootCmd.Flags().StringP("prefix", "p", "shot", "Filename prefix for batch processing")
	rootCmd.Flags().StringP("directory", "", ".", "Output directory for screenshots")

	// Add subcommands
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Show display information",
		Long:  `Display information about available screens and platform support. This command shows the number of displays, their resolutions, positions, and platform compatibility.`,
		Example: `  sshot info                    # Show all display information
  sshot info --verbose           # Show detailed information`,
		RunE: runInfo,
	}
	rootCmd.AddCommand(infoCmd)

	// Add usage tips
	rootCmd.SetHelpTemplate(`{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}

USAGE TIPS:
  • Use -r "x,y,width,height" for region capture
  • Use -c for clipboard-only (no file saved)
  • Use -n >1 for batch processing
  • Use -t with variables for dynamic filenames
  • Use --verbose for debugging information

COMMON PATTERNS:
  Quick capture:     sshot
  Region capture:    sshot -r "100,100,800,600"
  Clipboard only:    sshot -c
  Batch monitoring:  sshot -n 60 -i 60 -t "monitor_{time}.png"
  High quality:      sshot -f jpg -q 95 -o high_quality.jpg

For more examples, see: https://github.com/funnyzak/screenshot-cli`)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runScreenshot(cmd *cobra.Command, args []string) error {
	// Parse command line arguments
	config, err := config.ParseArgs(cmd, args)
	if err != nil {
		return fmt.Errorf("failed to parse arguments: %w", err)
	}

	if verbose {
		fmt.Printf("Configuration: %+v\n", config)
	}

	// Handle batch processing
	if config.Count > 1 {
		return batch.ProcessBatch(config)
	}

	// Single screenshot
	return captureSingleScreenshot(config)
}

func captureSingleScreenshot(config *config.Config) error {
	// Capture screenshot
	img, err := capture.CaptureScreen(config)
	if err != nil {
		return fmt.Errorf("failed to capture screenshot: %w", err)
	}

	// Process output
	if config.Clipboard {
		if err := output.CopyToClipboard(img); err != nil {
			return fmt.Errorf("failed to copy to clipboard: %w", err)
		}
		if verbose {
			fmt.Println("Screenshot copied to clipboard")
		}
	}

	// Save to file if output path is specified
	if config.OutputPath != "" {
		if err := output.SaveToFile(img, config); err != nil {
			return fmt.Errorf("failed to save file: %w", err)
		}
		if verbose {
			fmt.Printf("Screenshot saved to: %s\n", config.OutputPath)
		}
	}

	return nil
}

func runInfo(cmd *cobra.Command, args []string) error {
	// Check platform support
	if !capture.IsPlatformSupported() {
		fmt.Printf("Platform not supported: %s\n", capture.GetPlatformInfo())
		return nil
	}

	// Get display information
	displays, err := capture.GetDisplayInfo()
	if err != nil {
		return fmt.Errorf("failed to get display info: %w", err)
	}

	fmt.Printf("Platform: %s\n", capture.GetPlatformInfo())
	fmt.Printf("Active displays: %d\n\n", len(displays))

	for i, display := range displays {
		fmt.Printf("Display %d:\n", i)
		fmt.Printf("  Bounds: %s\n", display.String())
		fmt.Printf("  Size: %dx%d\n", display.Dx(), display.Dy())
		fmt.Printf("  Position: (%d, %d)\n", display.Min.X, display.Min.Y)
		fmt.Println()
	}

	return nil
}
