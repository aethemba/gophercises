package blackjack

import (
	"fmt"
	"gophercises/deck"
)

func main() {
	fmt.Println("Playing blackjack")
	cards := deck.New()
	_ = cards
}

type Player struct {
	Id      int
	Cards   []deck.Card
	isHouse bool
}

func (p Player) Total() int {
	return 0
}

type GameState struct {
	Players []Player
	Turn    int
	Deck    []deck.Card
}

func Play(p int) {
	var players []Player

	for i := 0; i < p; i++ {
		players = append(players, Player{isHouse: false, Id: i})
	}

	players = append(players, Player{isHouse: true, Id: p})

	cards := deck.New(deck.Deck(3))

	shuffledCards := deck.Shuffle(cards)

	gs := GameState{Players: players, Deck: shuffledCards}

	_ = gs

	for i := 0; i < 2; i++ {
		for _, player := range gs.Players {
			gs = DealCard(player, gs)
		}
	}

	for _, p := range gs.Players {
		if !p.isHouse {
			gs = PlayerTurn(gs)
		} else {
			gs = HouseTurn(gs)
		}

	}
}

func PlayerTurn(gs GameState) GameState {
	return gs
}

func HouseTurn(gs GameState) GameState {
	return gs
}

func DealCard(p Player, gs GameState) GameState {
	// Take top card from deck
	card, newDeck := gs.Deck[len(gs.Deck)-1], gs.Deck[:len(gs.Deck)-1]
	gs.Deck = newDeck

	// Assign it to player cards
	for _, player := range gs.Players {
		if player.Id == p.Id {
			player.Cards = append(player.Cards, card)
			break
		}
	}

	return gs
}

// Gameloop
//for {
// Check state for ending
// Action
//	  Who's move?
//	  What move?
//	  Do Move
// Update state
//}

// Determine players
//	- set player and house
//
