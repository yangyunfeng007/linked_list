package main

import (
	"fmt"
	"github.com/yangyunfeng007/linked_list"
)

func main() {
	l := linked_list.NewInt()
	for _, v := range []int{10, 12, 15} {
		if l.Insert(v) {
			fmt.Println("int list add", v)
		}
	}
	if l.Contains(10) {
		fmt.Println("int list contains 10")
	}
	l.Range(func(value int) bool {
		fmt.Println("int list find", value)
		return true
	})
	l.Delete(15)
	fmt.Printf("int list contains %d items\r\n", l.Len())
}
