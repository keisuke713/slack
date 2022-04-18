package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("user: ", os.Getenv("USER"))
}
