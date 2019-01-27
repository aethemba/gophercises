package main

import (
	"fmt"
	"gophercises/deck"
	"sort"
)

func main() {
	aceHearts := deck.Card{Value: deck.Ace, Suit: deck.Hearts}
	fmt.Println(aceHearts)

	cards := deck.New()

	deck.ShuffleInPlace(cards)

	sort.Sort(deck.ByNew(cards))

	for k, v := range cards {
		fmt.Println(k, v)
	}
}
