package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
)

func main() {
	fmt.Println(c.Green + "Hello, World!")
	fmt.Println(c.Cyan + "I am a computer!")
	fmt.Println(c.Red + "I am fast!" + c.X)
}
