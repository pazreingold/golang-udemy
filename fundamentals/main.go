package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
