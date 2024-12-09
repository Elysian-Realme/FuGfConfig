package main

import (
	"bufio"
	"fmt"
	"os"
)

// 流式写入文件 不附加
func writeFileWorker(filePath string, data <-chan string) error {
	// open file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for d := range data {
		_, err := file.WriteString(d)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeFile(filePath string, data []string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	for _, d := range data {
		fmt.Fprintln(write, d)
	}

	return write.Flush()
}
