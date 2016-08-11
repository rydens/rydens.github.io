package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)

	waffles := i.Ask(c.Green + "DO YOU LIKE WAFFLESLSLLS?" + c.Red + "\n> " + c.Cyan)
	if waffles == "yes" {
		if pancakes := i.Ask(c.G + "do you like the cakepans that make pancakes out of a pan that are cake?" + c.Red + "\n> " + c.Cyan); pancakes == "yes" {
			if frenchtoast := i.Ask(c.G + "YA LIKE THE TOAST OF FRANCE MUCH?" + c.Red + "\n> " + c.Cyan); frenchtoast == "yes" {
				fmt.Println(c.M + "DO DO DO CANTE WATE TOO GET A MOUTFUL (musiccnote)")
			} else {
				fmt.Println(c.Yellow + "WELL WE LIKE FRANCE'S TOAST")
			}
		} else {
			fmt.Println(c.Yellow + "well, I like the cake from pans made from pancakes from cakepans :(")
		}
	} else {
		fmt.Println(c.Yellow + "YOU SHOULD BE ASHAMMEMDED. I LIKE WAFFLESLSLLS. I ALSO SHOULD BE ASHAMMEMDED FOR MY BADE SPELINGG")
	}

}
