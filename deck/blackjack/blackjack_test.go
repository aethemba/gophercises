package blackjack

import (
	"gophercises/deck"
	"testing"
)

func TestDealCard(t *testing.T) {
	d := deck.New(deck.Deck(3))
	var players []Player
	players = append(players, Player{isHouse: false, Id: 1})

	shuffledCards := deck.Shuffle(d)
	gs := GameState{Players: players, Deck: shuffledCards}

	if len(gs.Deck) != 13*4*3 {
		t.Errorf("Insuffient number of deck cards. Expected %d, got %d", 13*4*3, len(gs.Deck))
	}

	if len(players[0].Cards) != 0 {
		t.Errorf("Insuffient number of cards for Player. Expected %d, got %d", 1, len(players[0].Cards))

	}

	gs = DealCard(players[0], gs)

	if len(gs.Deck) != (13*4*3)-1 {
		t.Errorf("Insuffient number of deck cards. Expected %d, got %d", 13*4*3, len(gs.Deck))
	}

}
