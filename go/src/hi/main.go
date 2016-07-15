package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
)

func main() {
	fmt.Println(c.Green + "Hi, what's your name?")
	name, err := i.Prompt(c.Red + "> " + c.Cyan)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Green + "Nice to meet you, " + name)
}
