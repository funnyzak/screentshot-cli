package config

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestParseRegion(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Region
		wantErr bool
	}{
		{
			name:    "valid region",
			input:   "100,200,800,600",
			want:    &Region{X: 100, Y: 200, Width: 800, Height: 600},
			wantErr: false,
		},
		{
			name:    "invalid format - too few parts",
			input:   "100,200,800",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid format - too many parts",
			input:   "100,200,800,600,900",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid coordinates - non-numeric",
			input:   "100,abc,800,600",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid region with spaces",
			input:   " 100 , 200 , 800 , 600 ",
			want:    &Region{X: 100, Y: 200, Width: 800, Height: 600},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRegion(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (got.X != tt.want.X || got.Y != tt.want.Y || got.Width != tt.want.Width || got.Height != tt.want.Height) {
				t.Errorf("parseRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidFormat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"png format", "png", true},
		{"jpg format", "jpg", true},
		{"jpeg format", "jpeg", true},
		{"bmp format", "bmp", true},
		{"gif format", "gif", true},
		{"uppercase PNG", "PNG", true},
		{"uppercase JPG", "JPG", true},
		{"invalid format", "tiff", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidFormat(tt.input); got != tt.want {
				t.Errorf("isValidFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().StringP("region", "r", "", "Region screenshot")
	cmd.Flags().StringP("output", "o", "screenshot.png", "Output file path")
	cmd.Flags().StringP("format", "f", "png", "Output format")
	cmd.Flags().IntP("quality", "q", 90, "JPG quality")
	cmd.Flags().BoolP("clipboard", "c", false, "Copy to clipboard")
	cmd.Flags().StringP("template", "t", "", "Filename template")
	cmd.Flags().IntP("count", "n", 1, "Number of screenshots")
	cmd.Flags().IntP("interval", "i", 1, "Screenshot interval")
	cmd.Flags().StringP("prefix", "p", "shot", "Filename prefix")
	cmd.Flags().StringP("dir", "d", ".", "Output directory")

	// Test valid arguments
	cmd.Flags().Set("region", "100,200,800,600")
	cmd.Flags().Set("format", "jpg")
	cmd.Flags().Set("quality", "85")

	config, err := ParseArgs(cmd, []string{"test.png"})
	if err != nil {
		t.Errorf("ParseArgs() error = %v", err)
		return
	}

	if config.OutputPath != "test.png" {
		t.Errorf("ParseArgs() output path = %v, want test.png", config.OutputPath)
	}

	if config.Region == nil || config.Region.X != 100 || config.Region.Y != 200 || config.Region.Width != 800 || config.Region.Height != 600 {
		t.Errorf("ParseArgs() region = %v, want {100, 200, 800, 600}", config.Region)
	}

	if config.Format != "jpg" {
		t.Errorf("ParseArgs() format = %v, want jpg", config.Format)
	}

	if config.Quality != 85 {
		t.Errorf("ParseArgs() quality = %v, want 85", config.Quality)
	}
}
