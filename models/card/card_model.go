package card

import (
	"Clue/models/player"
)

type Category int

const (
	Suspect Category = iota
	Place
	Weapon
)

type Card struct {
	Name     string
	WhoHasIt player.Player
	Category Category
}

var card1 = Card{Name: "Card1", WhoHasIt: player.Player{Id: 2, Name: "Name1"}, Category: Place}
var card2 = Card{Name: "Card2", WhoHasIt: player.Player{Id: 3, Name: "Name2"}, Category: Place}
var cards = []Card{card1, card2}

func GetAllCards() []Card {
	return cards
}

func GetCategories() map[Category]string {
	return map[Category]string{
		Suspect: "Suspect",
		Place:   "Place",
		Weapon:  "Weapon",
	}
}

func FillCards(filledCards []Card) []Card {
	cards = filledCards
	return cards
}

func AddCard(card Card) {
	cards = append(cards, card)
}

func RemoveAllCards() {
	cards = []Card{}
}
