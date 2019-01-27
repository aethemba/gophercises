package deck

import (
	"fmt"
	"math/rand"
)

type Value int
type Suit int

const (
	_ Value = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var Suits = []Suit{Diamonds, Hearts, Clubs, Spades}
var Values = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func Shuffle(d []Card) []Card {

	shuffled := make([]Card, len(d))

	for i, value := range rand.Perm(len(d)) {
		shuffled[value] = d[i]
	}

	return shuffled

}

type ByNew []Card

func (n ByNew) Len() int {
	return len(n)
}

func (n ByNew) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n ByNew) Less(i, j int) bool {

	// If i and j have the same value (e.g. both King),
	// than we look at the suit
	if n[i].Value == n[j].Value {
		return n[i].Suit < n[j].Suit
	}

	// I is less than J if it has a lower Value than J
	return n[i].Value < n[j].Value
}

func ShuffleInPlace(d []Card) {

	//For every card in the deck, swap with a random other card
	for i, _ := range d {
		otherCard := rand.Intn(len(d))
		d[otherCard], d[i] = d[i], d[otherCard]
	}
}

func New(options ...func(*[]Card)) []Card {
	cards := make([]Card, 0)

	for _, suit := range Suits {
		for _, val := range Values {
			cards = append(cards, Card{Suit: suit, Value: val})
		}
	}

	return cards
}

const (
	Spades Suit = iota
	Clubs
	Hearts
	Diamonds
)

type Card struct {
	Value Value
	Suit  Suit
}

func (c Card) String() string {
	var val, suit string
	switch c.Value {
	case 1:
		val = "Ace"
	case 2:
		val = "Two"
	case 3:
		val = "Three"
	case 4:
		val = "Four"
	case 5:
		val = "Five"
	case 6:
		val = "Six"
	case 7:
		val = "Seven"
	case 8:
		val = "Eight"
	case 9:
		val = "Nine"
	case 10:
		val = "Ten"
	case 11:
		val = "Jack"
	case 12:
		val = "Queen"
	case 13:
		val = "King"

	default:
		val = "Unknown value"
	}

	switch c.Suit {
	case 0:
		suit = "Spades"
	case 1:
		suit = "Clubs"
	case 2:
		suit = "Hearts"
	case 3:
		suit = "Diamonds"
	default:
		suit = "Unknown Suit"
	}

	return fmt.Sprintf("%s of %s", val, suit)
}
