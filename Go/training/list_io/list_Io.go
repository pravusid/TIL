package list_io

import (
	"bufio"
	"fmt"
	"io"
)

// WriteTo 함수는 문자열 슬라이스를 라인별로 출력함
func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

// ReadFrom 함수는 라인별 문자열을 읽음
func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
