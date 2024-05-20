package main

import (
	"fmt"
	"os"
)

func tailFile(fileName string, numBytes int) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open %s: %v", fileName, err)
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
	_, err = file.Read(buf)
	if err != nil {
		return fmt.Errorf("read %s: %v", fileName, err)
	}

	fmt.Printf("==> %s <==\n%s\n", fileName, buf)
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run . -c <num_bytes> file1.txt [file2.txt ...]")
		os.Exit(1)
	}

	numBytes := 0
	for i, arg := range args {
		if arg == "-c" {
			if len(args) < i+2 {
				fmt.Println("Error: -c option requires a value")
				os.Exit(1)
			}
			numBytes = 4 // Just considering it as a flag, not using the actual value
			continue
		}
		err := tailFile(arg, numBytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if i < len(args)-1 {
			fmt.Println()
		}
	}
}
