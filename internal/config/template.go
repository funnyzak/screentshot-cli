package config

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// TemplateProcessor handles filename template processing
type TemplateProcessor struct {
	counter int
}

// NewTemplateProcessor creates a new template processor
func NewTemplateProcessor() *TemplateProcessor {
	return &TemplateProcessor{
		counter: 1,
	}
}

// ProcessTemplate processes a filename template with variables
func (tp *TemplateProcessor) ProcessTemplate(template string, config *Config) string {
	if template == "" {
		return config.OutputPath
	}

	result := template

	// Replace template variables
	result = strings.ReplaceAll(result, "{timestamp}", strconv.FormatInt(time.Now().Unix(), 10))
	result = strings.ReplaceAll(result, "{datetime}", time.Now().Format("20060102_150405"))
	result = strings.ReplaceAll(result, "{date}", time.Now().Format("20060102"))
	result = strings.ReplaceAll(result, "{time}", time.Now().Format("150405"))
	result = strings.ReplaceAll(result, "{counter}", fmt.Sprintf("%03d", tp.counter))
	result = strings.ReplaceAll(result, "{random}", generateRandomString(6))
	result = strings.ReplaceAll(result, "{prefix}", config.Prefix)

	// Add file extension if not present
	if !hasFileExtension(result) {
		result += "." + config.Format
	}

	return result
}

// IncrementCounter increments the internal counter
func (tp *TemplateProcessor) IncrementCounter() {
	tp.counter++
}

// SetCounter sets the counter to a specific value
func (tp *TemplateProcessor) SetCounter(value int) {
	tp.counter = value
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// hasFileExtension checks if a filename has a file extension
func hasFileExtension(filename string) bool {
	extensions := []string{".png", ".jpg", ".jpeg", ".bmp", ".gif"}
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}
	return false
}
