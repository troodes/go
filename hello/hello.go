package main

// https://github.com/GoClipse/goclipse/blob/latest/documentation/UserGuide.md#user-guide

import (
	"fmt"
	"github.com/troodes/sandbox/stringutil"
)

func main() {
    fmt.Printf("hello, world.\n")
    fmt.Printf(stringutil.Reverse("hello, world\n"))
}
