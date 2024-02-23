package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// 置換対象のファイル名
	filename := "client.go"

	// ファイルを読み込みモードで開く
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Sprintf("ファイルを開けませんでした: %v", err))
	}
	defer file.Close()

	// ファイルの内容を読み込む
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		// 文字列の置換を行う
		target := regexp.MustCompile("[(]time.RFC3339[)]")
		line = target.ReplaceAllString(line, "(time.RFC3339Nano)")
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("ファイルを読み込めませんでした: %v", err))
	}

	// ファイルの内容を一度クリア
	file.Truncate(0)
	file.Seek(0, 0)

	// 置換した内容をファイルに書き込む
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(fmt.Sprintf("ファイルに書き込めませんでした: %v", err))
		}
	}
	writer.Flush()

	fmt.Println("ファイルの置換が完了しました:", filename)
}
