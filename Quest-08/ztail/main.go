package main

import (
	"fmt"
	"os"
)

func parseNumChars(numChars string) int {
	num := 0
	for _, char := range numChars {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			break
		}
	}
	return num
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func tailFile(filename, numChars string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open %s: %v", filename, err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat %s: %v", filename, err)
	}

	fileSize := fileStat.Size()
	num := parseNumChars(numChars)

	if fileSize < int64(num) {
		num = int(fileSize)
	}

	buffer := make([]byte, num)
	_, err = file.ReadAt(buffer, fileSize-int64(num))
	if err != nil {
		return fmt.Errorf("read %s: %v", filename, err)
	}

	fmt.Print(string(buffer))
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . -c <num> file1.txt [file2.txt ...]")
		os.Exit(1)
	}

	numChars := os.Args[2]
	filenames := os.Args[3:]

	exitStatus := 0

	for i, filename := range filenames {
		if fileExists(filename) {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", filename)
			if err := tailFile(filename, numChars); err != nil {
				fmt.Println(err)
				exitStatus = 1
			}
		} else {
			fmt.Fprintf(os.Stderr, "open %s: no such file or directory\n", filename)
			exitStatus = 1
		}
	}

	os.Exit(exitStatus)
}
