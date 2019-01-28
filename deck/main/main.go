package main

import (
	"fmt"
	"gophercises/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))

	for i := range h {
		strs[i] = h[i].String()
	}
	return fmt.Sprintf("%s", strings.Join(strs, ", "))
}

func (h Hand) DealerString() string {
	return h[0].String() + ", ** HIDDEN **"
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func main() {
	fmt.Println("Playing blackjack")
	cards := deck.New(deck.Deck(3))
	cards = deck.Shuffle(cards)
	var card deck.Card

	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, h := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*h = append(*h, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Hand player: ", player)
		fmt.Println("Hand dealer: ", dealer.DealerString())
		fmt.Println("Action? (H)it or (S)tand?")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	fmt.Println("**Final hands**")
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)

}
