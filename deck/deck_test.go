package deck

import "fmt"

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
