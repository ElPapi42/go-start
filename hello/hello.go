package main

import (
	"fmt"

	"github.com/elpapi42/go-start/phone"
	"github.com/elpapi42/go-start/practice"
	"github.com/elpapi42/go-start/strutil"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf(strutil.Reverse("hello"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))

	mia3 := phone.New("Mi A3", "Xiaomi", 6)
	mia3.Features()

	practice.CreateVars()
}
