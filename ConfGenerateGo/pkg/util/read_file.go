package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadFile 读取文件 忽略空行及注释
func ReadFile(filePath string) []string {
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
		if !IsNote(str) {
			str = FormatCorrection(str)
			ans = append(ans, str)
		}
	}
	return ans
}