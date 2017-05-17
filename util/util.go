package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadLines(file string) ([]string, error) {
	return ReadSplitter(file, '\n')
}

func ReadSplitter(file string, splitter byte) (lines []string, err error) {
	fin, err := os.Open(file)
	if err != nil {
		return
	}

	r := bufio.NewReader(fin)
	for {
		line, err := r.ReadString(splitter)
		if err == io.EOF {
			break
		}
		line = strings.Replace(line, string(splitter), "", -1)
		lines = append(lines, line)
	}
	return
}
