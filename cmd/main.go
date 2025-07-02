package main

import (
	"fmt"
	"os"

	"screenshot-cli/internal/batch"
	"screenshot-cli/internal/capture"
	"screenshot-cli/internal/config"
	"screenshot-cli/internal/output"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	verbose bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "sshot [flags] [output_file]",
		Short:   "A lightweight screenshot CLI tool",
		Long:    `A lightweight, dependency-free Windows command line screenshot tool that supports full screen and region screenshots with flexible output control and batch processing capabilities.`,
		Version: version,
		RunE:    runScreenshot,
	}

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "Verbose output mode")

	// Screenshot flags
	rootCmd.Flags().StringP("region", "r", "", "Region screenshot \"x,y,width,height\"")
	rootCmd.Flags().StringP("output", "o", "screenshot.png", "Output file path")

	// Output control flags
	rootCmd.Flags().StringP("format", "f", "png", "Output format (png/jpg/bmp/gif)")
	rootCmd.Flags().IntP("quality", "q", 90, "JPG compression quality (1-100)")
	rootCmd.Flags().BoolP("clipboard", "c", false, "Copy to clipboard")
	rootCmd.Flags().StringP("template", "t", "", "Filename template")

	// Batch processing flags
	rootCmd.Flags().IntP("count", "n", 1, "Number of screenshots")
	rootCmd.Flags().IntP("interval", "i", 1, "Screenshot interval (seconds)")
	rootCmd.Flags().StringP("prefix", "p", "shot", "Filename prefix")
	rootCmd.Flags().StringP("dir", "d", ".", "Output directory")

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
