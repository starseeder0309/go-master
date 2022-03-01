package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
		fmt.Println("[EX] ex26-01 word filePath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files)
}

func GetFiles(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		targetFiles, err := GetFiles(path)
		if err != nil {
			fmt.Println("잘못된 파일 경로입니다. err:", err, "path:", path)
			return
		}

		fmt.Println("찾으려는 파일 목록")
		for _, name := range targetFiles {
			fmt.Println(name)
		}
	}
}
