package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// readFile 读取文件 忽略空行及注释
func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var ans []string
	for {
		str, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if !isNote(str) {
			str = formatCorrection(str)
			ans = append(ans, str)
		}
	}
	return ans
}
