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
		State:  gs.State,
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

func Shuffle(gs GameState) GameState {
	ret := clone(gs)

	cards := deck.New(deck.Deck(3))
	ret.Deck = deck.Shuffle(cards)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)

	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)

		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}

	ret.State = StatePlayerTurn

	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)

	hand := ret.CurrentPlayer()
	fmt.Println("Hit", hand)
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)

	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)

	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()

	fmt.Println("**Final hands**")
	fmt.Println("Player: ", ret.Player, "\nScore: ", pScore)
	fmt.Println("Dealer: ", ret.Dealer, "\nScore: ", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose...")
	case dScore == pScore:
		fmt.Println("Draw!")
	}

	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)

	gs = Deal(gs)

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Hand player: ", gs.Player)
		fmt.Println("Hand dealer: ", gs.Dealer.DealerString())
		fmt.Println("Action? (H)it or (S)tand?")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid option. Choose (h)it or (s)tand")
		}

	}

	for gs.State == StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() < 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndHand(gs)
}
