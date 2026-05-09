package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func main() {
	// Get the current working directory
	searchDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scanning directory: %s\n", searchDir)

	files, err := os.ReadDir(searchDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			processImage(file.Name())
		}
	}
	
	fmt.Println("Done! Press Enter to exit.")
	fmt.Scanln() 
}

func processImage(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", fileName, err)
		return
	}
	// We decode the image into memory
	img, _, err := image.Decode(file)
	file.Close() // Close original immediately after decoding

	if err != nil {
		fmt.Printf("Error decoding %s: %v\n", fileName, err)
		return
	}

	outputName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".webp"
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", outputName, err)
		return
	}

	// Encode with 75% quality
	err = webp.Encode(outFile, img, &webp.Options{Lossless: false, Quality: 75})
	outFile.Close()

	if err != nil {
		fmt.Printf("Error encoding %s: %v\n", outputName, err)
		return
	}

	// Delete the original Windows file
	err = os.Remove(fileName)
	if err != nil {
		fmt.Printf("Converted %s but couldn't delete original: %v\n", outputName, err)
	} else {
		fmt.Printf("Converted and Deleted: %s\n", fileName)
	}
}