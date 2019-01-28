package main

import (
	"fmt"
	"gophercises/deck"
)

func main() {
	aceHearts := deck.Card{Value: deck.Ace, Suit: deck.Hearts}
	fmt.Println(aceHearts)

	cards := deck.New()

	deck.ShuffleInPlace(cards)

	//sort.Sort(deck.ByNew(cards))

	for k, v := range cards {
		fmt.Println(k, v)
	}

	fmt.Println("***")

	newCards := deck.New(deck.DefaultSort, deck.Jokers(3))
	for k, v := range newCards {
		fmt.Println(k, v)
	}
}

type Player struct {
	Hand []deck.Card
}

func DealCards(amount int) {

}

// BLACKJACK
// Each player gets two cards (players both visible, house 1 visible)
