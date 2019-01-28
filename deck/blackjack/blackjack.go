package blackjack

import (
	"fmt"
	"gophercises/deck"
)

func main() {
	fmt.Println("Playing blackjack")
	cards := deck.New(deck.Deck(3))
	cards = deck.Shuffle(cards)

	for i := 0; i < 10; i++ {
		var card deck.Card
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}

}

type Player struct {
	Id      int
	Cards   []deck.Card
	isHouse bool
}

func (p Player) Total() int {
	return 0
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

	for i := 0; i < 2; i++ {
		for _, player := range gs.Players {
			gs = DealCard(player, gs)
		}
	}

	for _, p := range gs.Players {
		if !p.isHouse {
			gs = PlayerTurn(p, gs)
		} else {
			gs = HouseTurn(p, gs)
		}
	}

	winner := DetermineWinner(gs)
	fmt.Printf("Winner is %#v\n", winner)
}

func DetermineWinner(gs GameState) Player {
	for _, player := range gs.Players {
		_ = player
	}
	return gs.Players[0]
}

func PlayerTurn(p Player, gs GameState) GameState {
	fmt.Println("You have the following cards")
	return gs
}

func HouseTurn(p Player, gs GameState) GameState {
	return gs
}

func Draw(p Player, gs GameState) GameState {
	// Take top card from deck
	var card deck.Card
	card, gs.Deck = gs.Deck[0], gs.Deck[1:]

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
