package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("user iss", os.Getenv("USER"))
}
