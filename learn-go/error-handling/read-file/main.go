package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("temp.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := strings.NewReader("Hello test")

	io.Copy(f, r)
}
