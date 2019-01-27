package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

const (
	MinValue = Ace
	MaxValue = King
)

var Suits = []Suit{Diamonds, Hearts, Clubs, Spades}

// var Values = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func Shuffle(d []Card) []Card {

	shuffled := make([]Card, len(d))

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i, value := range r.Perm(len(d)) {
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

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Jokers(j int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 1; i <= j; i++ {
			cards = append(cards, Card{Value: 0, Suit: Joker})
		}
		return cards
	}
}

func Filter(unwanted []Card) func([]Card) []Card {
	return func(cards []Card) []Card {
		for _, unwantedCard := range unwanted {
			for i, j := range cards {
				if j.Suit == unwantedCard.Suit && j.Value == unwantedCard.Value {
					cards = append(cards[:i], cards[i+1:]...)
				}
			}
		}
		return cards
	}
}

// Alternative func
func FilterFunc(f func(Card) bool) func([]Card) []Card {
	// f returns true if the card should be filtered
	return func(cards []Card) []Card {
		var result []Card
		for _, card := range cards {
			if !f(card) {
				result = append(result, card)
			}
		}
		return result
	}
}

func CustomSort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		var bn ByNew = cards
		return bn.Less(i, j)
	}
}

func Deck(d int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card

		for i := 0; i < d; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

func ShuffleInPlace(d []Card) {

	//For every card in the deck, swap with a random other card
	for i, _ := range d {
		otherCard := rand.Intn(len(d))
		d[otherCard], d[i] = d[i], d[otherCard]
	}
}

func New(options ...func([]Card) []Card) []Card {
	cards := make([]Card, 0)

	for _, suit := range Suits {
		for val := MinValue; val <= MaxValue; val++ {
			cards = append(cards, Card{Suit: suit, Value: val})
		}
	}

	for _, opt := range options {
		cards = opt(cards)
	}

	return cards
}

const (
	Spades Suit = iota
	Clubs
	Hearts
	Diamonds
	Joker
)

type Card struct {
	Value Value
	Suit  Suit
}

func (c Card) String() string {
	var val, suit string
	switch c.Value {
	case 0:
		val = "Th"
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
	case 4:
		suit = "Joker"
	default:
		suit = "Unknown Suit"
	}

	if c.Suit == Joker {
		return fmt.Sprintf("A Joker")
	}

	return fmt.Sprintf("%s of %s", val, suit)
}
