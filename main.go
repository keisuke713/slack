package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("user is", os.Getenv("USER"))
}
