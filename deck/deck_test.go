package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Value: Two, Suit: Hearts})
	fmt.Println(Card{Value: Ace, Suit: Spades})
	fmt.Println(Card{Value: King, Suit: Clubs})
	fmt.Println(Card{Value: Nine, Suit: Diamonds})

	// Output:
	// Two of Hearts
	// Ace of Spades
	// King of Clubs
	// Nine of Diamonds
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Invalid number of cards")
	}
}

func TestJokers(t *testing.T) {
	jokerCount := 3
	cards := New(DefaultSort, Jokers(jokerCount))

	for i := 1; i <= jokerCount; i++ {
		if cards[len(cards)-i].Suit != Joker {
			t.Error("Expected a Joker, received ", cards[len(cards)])
		}
	}

}

func TestFilter(t *testing.T) {
	unwantedCards := []Card{Card{Suit: Spades, Value: Ace}, Card{Suit: Diamonds, Value: Three}}

	cards := New(DefaultSort, Filter(unwantedCards))

	for _, card := range cards {
		if card.Suit == Spades && card.Value == Ace {
			t.Error("Found Ace of Spades that should be filtered")
		}

		if card.Suit == Diamonds && card.Value == Three {
			t.Error("Found Three of Diamonds that should be filtered")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	// 13 cards of each suite, times three decks
	if len(cards) != 13*4*3 {
		t.Errorf("Incorrect number of cards. Expected %d, received %d", 13*4*3, len(cards))
	}

}
