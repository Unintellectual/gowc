package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func countWords(filename string) (int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a scanner to read the file word by word
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Count the words
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func countBytes(filename string) (int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Get file information to determine file size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Return the file size in bytes
	return int(fileInfo.Size()), nil
}

func main() {
	// Define and parse flags
	filename := flag.String("file", "", "Input file name")
	countWordsFlag := flag.Bool("w", false, "Display the word count")
	countBytesFlag := flag.Bool("c", false, "Display the byte count")
	flag.Parse()

	// Check if filename flag is provided
	if *filename == "" {
		fmt.Println("Usage: go run main.go -file <filename> [-w] [-c]")
		os.Exit(1)
	}

	// Display the word count if -w flag is provided
	if *countWordsFlag {
		wordCount, err := countWords(*filename)
		if err != nil {
			fmt.Printf("Error counting words: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Word count: %d\n", wordCount)
	}

	// Display the byte count if -c flag is provided
	if *countBytesFlag {
		byteCount, err := countBytes(*filename)
		if err != nil {
			fmt.Printf("Error counting bytes: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Byte count: %d\n", byteCount)
	}

	// If neither flag is provided, display an error message
	if !*countWordsFlag && !*countBytesFlag {
		fmt.Println("Please specify either -w or -c flag")
		os.Exit(1)
	}
}
