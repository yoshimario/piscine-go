package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printFileContent(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		errorMessage := "ERROR: open " + fileName + ": no such file or directory\n"
		for _, char := range errorMessage {
			z01.PrintRune(char)
		}
		return err
	}
	for _, char := range string(content) {
		z01.PrintRune(char)
	}
	return nil
}

func printStdin() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		for _, char := range string(buf[:n]) {
			z01.PrintRune(char)
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printStdin()
	} else {
		for _, fileName := range args {
			if err := printFileContent(fileName); err != nil {
				os.Exit(1)
			}
		}
	}
}
