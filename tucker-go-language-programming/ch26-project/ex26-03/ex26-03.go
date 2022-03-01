package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNumber int
	conent     string
}

type FileInfo struct {
	fileName string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니.")
		fmt.Println("ex26-03 word filePath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	foundFiles := []FileInfo{}

	for _, path := range files {
		foundFiles = append(foundFiles, FindWordInAllFiles(word, path)...)
	}

	for _, foundFile := range foundFiles {
		fmt.Println(foundFile.fileName)
		fmt.Println("=========================")

		for _, line := range foundFile.lines {
			fmt.Println("\t", line.lineNumber, "\t", line.conent)
		}

		fmt.Println("=========================")
		fmt.Println()
	}
}

func GetFiles(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FileInfo {
	foundFiles := []FileInfo{}

	files, err := GetFiles(path)
	if err != nil {
		fmt.Println("잘못된 파일 경로입니다. err:", err, "path:", path)
		return foundFiles
	}

	for _, fileName := range files {
		foundFiles = append(foundFiles, FindWordInFile(word, fileName))
	}

	return foundFiles
}

func FindWordInFile(word, fileName string) FileInfo {
	foundFile := FileInfo{fileName, []LineInfo{}}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", fileName)
		return foundFile
	}

	defer file.Close()

	lineNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			foundFile.lines = append(foundFile.lines, LineInfo{lineNumber, line})
		}
		lineNumber++
	}

	return foundFile
}
