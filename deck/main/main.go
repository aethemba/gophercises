package main

import (
	"fmt"
	"gophercises/deck"
	"strings"
)

type Hand []deck.Card

type State int

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
	Turn   int
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("Illegal state")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
		Turn:   gs.Turn,
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Value), 10)
	}
	return score
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Value == deck.Ace {
			// Ace is currently already 1. Can only use 1 Ace
			return minScore + 10
		}
	}
	return minScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
	var gs GameState
	gs.Deck = deck.New(deck.Deck(3))
	// fmt.Println("Playing blackjack")
	// cards := deck.New(deck.Deck(3))
	// cards = deck.Shuffle(cards)
	// var card deck.Card

	// var player, dealer Hand
	// for i := 0; i < 2; i++ {
	// 	for _, h := range []*Hand{&player, &dealer} {
	// 		card, cards = draw(cards)
	// 		*h = append(*h, card)
	// 	}
	// }

	// var input string
	// for input != "s" {
	// 	fmt.Println("Hand player: ", player)
	// 	fmt.Println("Hand dealer: ", dealer.DealerString())
	// 	fmt.Println("Action? (H)it or (S)tand?")
	// 	fmt.Scanf("%s\n", &input)

	// 	switch input {
	// 	case "h":
	// 		card, cards = draw(cards)
	// 		player = append(player, card)
	// 	}
	// }

	// if dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() < 17) {
	// 	card, cards = draw(cards)
	// 	dealer = append(dealer, card)
	// }

	// pScore, dScore := player.Score(), dealer.Score()

	// fmt.Println("**Final hands**")
	// fmt.Println("Player: ", player, "\nScore: ", pScore)
	// fmt.Println("Dealer: ", dealer, "\nScore: ", dScore)

	// switch {
	// case pScore > 21:
	// 	fmt.Println("You busted")
	// case dScore > 21:
	// 	fmt.Println("Dealer busted")
	// case pScore > dScore:
	// 	fmt.Println("You win!")
	// case dScore > pScore:
	// 	fmt.Println("You lose...")
	// case dScore == pScore:
	// 	fmt.Println("Draw!")
	// }
}
