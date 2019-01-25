package deck

import "fmt"

type Value int
type Suit int

const (
	Ace Value = iota
	One
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
var Values = []Value{Ace, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func Shuffle(*[]Card) {
	fmt.Println("Making some random permutation on the cards")
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
	Diamonds Suit = iota
	Hearts
	Clubs
	Spades
)

type Card struct {
	Value Value
	Suit  Suit
}

func (c Card) String() string {
	var val, suit string
	switch c.Value {
	case 0:
		val = "Ace"
	case 1:
		val = "One"
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
		val = "Unknown card value"
	}

	switch c.Suit {
	case 0:
		suit = "Diamonds"
	case 1:
		suit = "Hearts"
	case 2:
		suit = "Clubs"
	case 3:
		suit = "Spades"
	default:
		suit = "Unknown Suit"
	}

	return fmt.Sprintf("%s of %s", val, suit)
}
