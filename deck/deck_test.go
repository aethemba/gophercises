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
