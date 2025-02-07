package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello , world")
	fmt.Println(morestrings.ReverseRunes("hello , world"))
	fmt.Println(cmp.Diff("hello", "hello go"))

}
