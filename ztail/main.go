package main

import (
	"fmt"
	"os"
)

func tailFile(fileName string, numBytes int) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open %s: no such file or directory", fileName)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat %s: %v", fileName, err)
	}

	offset := stat.Size() - int64(numBytes)
	if offset < 0 {
		offset = 0
	}

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("seek %s: %v", fileName, err)
	}

	buf := make([]byte, numBytes)
	n, err := file.Read(buf)
	if err != nil {
		return fmt.Errorf("read %s: %v", fileName, err)
	}

	fmt.Printf("==> %s <==\n", fileName)
	fmt.Print(string(buf[:n]))
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <num_bytes> file1.txt [file2.txt ...]")
		os.Exit(1)
	}

	numBytes := 0
	if _, err := fmt.Sscanf(args[1], "%d", &numBytes); err != nil || numBytes <= 0 {
		fmt.Println("Error: invalid number of bytes")
		os.Exit(1)
	}

	exitCode := 0
	errorsPrinted := 0

	for i, fileName := range args[2:] {
		err := tailFile(fileName, numBytes)
		if err != nil {
			fmt.Println(err)
			errorsPrinted++
			if errorsPrinted == 2 && i < len(args)-3 {
				fmt.Println()
			}
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
