package batch

import (
	"fmt"
	"path/filepath"
	"time"

	"screenshot-cli/internal/capture"
	"screenshot-cli/internal/config"
	"screenshot-cli/internal/output"
)

// ProcessBatch handles batch screenshot processing
func ProcessBatch(cfg *config.Config) error {
	templateProcessor := cfg.NewTemplateProcessor()

	fmt.Printf("Starting batch capture: %d screenshots, %d second intervals\n", cfg.Count, cfg.Interval)

	for i := 1; i <= cfg.Count; i++ {
		// Update counter for template processing
		templateProcessor.SetCounter(i)

		// Capture screenshot
		img, err := capture.CaptureScreen(cfg)
		if err != nil {
			return fmt.Errorf("failed to capture screenshot %d: %w", i, err)
		}

		// Generate output path
		outputPath := output.GetOutputPath(cfg, templateProcessor)

		// Ensure full path includes directory
		if cfg.Dir != "." {
			outputPath = filepath.Join(cfg.Dir, filepath.Base(outputPath))
		}

		// Create temporary config for saving
		saveConfig := &config.Config{
			OutputPath: outputPath,
			Format:     cfg.Format,
			Quality:    cfg.Quality,
		}

		// Save to file
		if err := output.SaveToFile(img, saveConfig); err != nil {
			return fmt.Errorf("failed to save screenshot %d: %w", i, err)
		}

		fmt.Printf("Screenshot %d/%d saved: %s\n", i, cfg.Count, outputPath)

		// Copy to clipboard if requested
		if cfg.Clipboard {
			if err := output.CopyToClipboard(img); err != nil {
				fmt.Printf("Warning: failed to copy screenshot %d to clipboard: %v\n", i, err)
			}
		}

		// Wait before next screenshot (except for the last one)
		if i < cfg.Count {
			time.Sleep(time.Duration(cfg.Interval) * time.Second)
		}
	}

	fmt.Printf("Batch capture completed: %d screenshots saved\n", cfg.Count)
	return nil
}

// ProcessBatchWithProgress handles batch processing with detailed progress information
func ProcessBatchWithProgress(cfg *config.Config) error {
	templateProcessor := cfg.NewTemplateProcessor()

	fmt.Printf("Starting batch capture: %d screenshots, %d second intervals\n", cfg.Count, cfg.Interval)
	fmt.Printf("Output directory: %s\n", cfg.Dir)
	fmt.Printf("Format: %s, Quality: %d\n", cfg.Format, cfg.Quality)

	startTime := time.Now()

	for i := 1; i <= cfg.Count; i++ {
		iterationStart := time.Now()

		// Update counter for template processing
		templateProcessor.SetCounter(i)

		// Capture screenshot
		img, err := capture.CaptureScreen(cfg)
		if err != nil {
			return fmt.Errorf("failed to capture screenshot %d: %w", i, err)
		}

		// Generate output path
		outputPath := output.GetOutputPath(cfg, templateProcessor)

		// Ensure full path includes directory
		if cfg.Dir != "." {
			outputPath = filepath.Join(cfg.Dir, filepath.Base(outputPath))
		}

		// Create temporary config for saving
		saveConfig := &config.Config{
			OutputPath: outputPath,
			Format:     cfg.Format,
			Quality:    cfg.Quality,
		}

		// Save to file
		if err := output.SaveToFile(img, saveConfig); err != nil {
			return fmt.Errorf("failed to save screenshot %d: %w", i, err)
		}

		// Get image info for progress
		width, height, _ := output.GetImageInfo(img)
		iterationDuration := time.Since(iterationStart)

		fmt.Printf("[%d/%d] Saved: %s (%dx%d) - %v\n",
			i, cfg.Count, outputPath, width, height, iterationDuration)

		// Copy to clipboard if requested
		if cfg.Clipboard {
			if err := output.CopyToClipboard(img); err != nil {
				fmt.Printf("Warning: failed to copy screenshot %d to clipboard: %v\n", i, err)
			}
		}

		// Wait before next screenshot (except for the last one)
		if i < cfg.Count {
			time.Sleep(time.Duration(cfg.Interval) * time.Second)
		}
	}

	totalDuration := time.Since(startTime)
	fmt.Printf("Batch capture completed: %d screenshots in %v\n", cfg.Count, totalDuration)

	return nil
}
