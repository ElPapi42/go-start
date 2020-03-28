package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/elpapi42/go-start/phone"
	"github.com/elpapi42/go-start/strutil"
)

func main() {
	fmt.Printf(strutil.Reverse("hello"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))

	mia3 := phone.New("Mi A3", "Xiaomi", 6)
	mia3.Features()
}
