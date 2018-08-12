package hof

import (
	"bufio"
	"io"
)

// ReadFrom 함수는 인자로 받은 함수의 방식대로 처리한 데이터를 한 줄씩 읽는다
func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
