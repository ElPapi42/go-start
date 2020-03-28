package main

import (
	"fmt"

	"github.com/elpapi42/go-start/strutil"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf(strutil.Reverse("hello"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))
}
