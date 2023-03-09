package main

import (
	"fmt"
	"geometry"
	"github.com/hackebrot/turtle"
	"github.com/swam-inc/stringmod/quotes"
	"github.com/swam-inc/stringmod/strings"
)

func main() {
	emoji, ok := turtle.Emojis["turtle"]
	if !ok {
		fmt.Println("No emoji found")
	} else {
		fmt.Println(emoji.Char)
	}
	fmt.Print("the main() program under package main start...")
	pt1 := geometry.Point{X: 2, Y: 3}
	fmt.Print(pt1)
	fmt.Println(pt1.Length())

	odd, even := strings.CountOddEven("12345")
	fmt.Println(odd, even)

	fmt.Println(quotes.GetEmoji("turtle"))
}
