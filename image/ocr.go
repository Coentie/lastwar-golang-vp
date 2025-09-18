package image

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"lastwar/notifier/hdadb"
	"os"
	"os/exec"
	"strings"
)

const startX = 173
const startY = 248
const startWidth = 309
const startHeight = 55

func ReadNameFromRegion() (string, error) {
	hdadb.PrintScreen()
	screenshot, err := OpenImage("screen.png")
	if err != nil {
		panic(err)
	}

	cropped := GetCrop(startX, startY, startWidth, startHeight, screenshot)
	err = saveImage(cropped, "cropped_region.png")
	if err != nil {
		panic(err)
	}

	text, err := OCRWithTesseract("cropped_region.png")

	if err != nil {
		fmt.Println("Could not read name")
		return "", err
	}

	return text, nil
}

// saveImage saves an image to a file
func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode as PNG and save
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func OCRWithTesseract(imagePath string) (string, error) {
	// Check if image file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return "", fmt.Errorf("Image file not found: %s", imagePath)
	}

	// Create a temporary output file (Tesseract will add .txt extension)
	tempOutput := "temp_ocr_output"
	defer os.Remove(tempOutput + ".txt") // Clean up temporary file

	// Build the command
	cmd := exec.CommandContext(context.Background(),
		os.Getenv("OCR_PATH"),
		imagePath,
		tempOutput,
		"--psm", "7",
		"-c", fmt.Sprintf("tessedit_char_whitelist=%s", os.Getenv("OCR_CHAR_SET")))

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Tesseract execution failed: %v, output: %s", err, string(output))
	}

	// Read the output file
	outputFile := tempOutput + ".txt"
	content, err := os.ReadFile(outputFile)
	if err != nil {
		return "", fmt.Errorf("Failed to read OCR output file: %v", err)
	}

	// Clean up the text
	text := strings.TrimSpace(string(content))
	return text, nil
}
