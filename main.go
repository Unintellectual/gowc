package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	// Create a scanner to read the file word by word
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Count the words
	charCount := 0
	for scanner.Scan() {
		charCount++
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return charCount, nil
}
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
	countWordsFlag := flag.Bool("w", false, "Display the word count")
	countBytesFlag := flag.Bool("c", false, "Display the byte count")
	countLinesFlag := flag.Bool("m", false, "Display the amount of lines in file")
	flag.Parse()

	// Get the filename from command line arguments
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run wordcount.go [-w] [-c] [-m] <filename>")
		os.Exit(1)
	}
	filename := args[0]

	// Display the word count if -w flag is provided
	if *countWordsFlag {
		wordCount, err := countWords(filename)
		if err != nil {
			fmt.Printf("Error counting words: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Word count: %d\n", wordCount)
	}
	// Display the word count if -m flag is provided
	if *countLinesFlag {
		charCount, err := countLines(filename)
		if err != nil {
			fmt.Printf("Error counting characters in file: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Word count: %d\n", charCount)
	}

	// Display the byte count if -c flag is provided
	if *countBytesFlag {
		byteCount, err := countBytes(filename)
		if err != nil {
			fmt.Printf("Error counting bytes: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Byte count: %d\n", byteCount)
	}

	// If neither flag is provided, display an error message
	if !*countWordsFlag && !*countBytesFlag && !*countLinesFlag {
		fmt.Println("Please specify either -w or -c or -l flag")
		os.Exit(1)
	}
}
