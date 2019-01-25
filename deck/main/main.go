package main

import (
	"fmt"
	"gophercises/deck"
)

func main() {
	aceHearts := deck.Card{Value: deck.Ace, Suit: deck.Hearts}
	fmt.Println(aceHearts)

	cards := deck.New()

	fmt.Println("Cards in deck")
	for _, c := range cards {
		fmt.Println(c)
	}
}
