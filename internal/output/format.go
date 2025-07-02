package output

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"

	"golang.org/x/image/bmp"
)

// ImageFormat represents supported image formats
type ImageFormat string

const (
	FormatPNG  ImageFormat = "png"
	FormatJPG  ImageFormat = "jpg"
	FormatJPEG ImageFormat = "jpeg"
	FormatBMP  ImageFormat = "bmp"
	FormatGIF  ImageFormat = "gif"
)

// EncodeImage encodes an image to the specified format with quality settings
func EncodeImage(img image.Image, format ImageFormat, quality int) ([]byte, error) {
	var buf bytes.Buffer

	switch strings.ToLower(string(format)) {
	case string(FormatPNG):
		if err := png.Encode(&buf, img); err != nil {
			return nil, fmt.Errorf("failed to encode PNG: %w", err)
		}

	case string(FormatJPG), string(FormatJPEG):
		opts := jpeg.Options{Quality: quality}
		if err := jpeg.Encode(&buf, img, &opts); err != nil {
			return nil, fmt.Errorf("failed to encode JPEG: %w", err)
		}

	case string(FormatBMP):
		if err := bmp.Encode(&buf, img); err != nil {
			return nil, fmt.Errorf("failed to encode BMP: %w", err)
		}

	case string(FormatGIF):
		if err := gif.Encode(&buf, img, nil); err != nil {
			return nil, fmt.Errorf("failed to encode GIF: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return buf.Bytes(), nil
}

// GetFileExtension returns the file extension for a given format
func GetFileExtension(format ImageFormat) string {
	switch strings.ToLower(string(format)) {
	case string(FormatPNG):
		return ".png"
	case string(FormatJPG), string(FormatJPEG):
		return ".jpg"
	case string(FormatBMP):
		return ".bmp"
	case string(FormatGIF):
		return ".gif"
	default:
		return ".png"
	}
}

// IsFormatSupported checks if a format is supported
func IsFormatSupported(format string) bool {
	supportedFormats := map[string]bool{
		"png":  true,
		"jpg":  true,
		"jpeg": true,
		"bmp":  true,
		"gif":  true,
	}
	return supportedFormats[strings.ToLower(format)]
}

// GetImageInfo returns basic information about an image
func GetImageInfo(img image.Image) (width, height int, format string) {
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy(), "image"
}
