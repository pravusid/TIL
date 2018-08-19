package sort

import (
	"container/heap"
	"fmt"
)

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}
