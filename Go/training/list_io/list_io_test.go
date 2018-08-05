package list_io

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string{
		"kim@korea.kr",
		"park@korea.kr",
		"lee@korea.kr",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// kim@korea.kr
	// park@korea.kr
	// lee@korea.kr
}

func ExampleReadFrom() {
	r := strings.NewReader("kim\npark\nlee\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [kim park lee]
}
