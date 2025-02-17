package card

import (
	"reflect"
	"testing"
)

func TestGetAllCards(t *testing.T) {
	card1 := Card{Name: "Card1"}
	card2 := Card{Name: "Card2"}
	cards = []Card{card1, card2}

	result := GetAllCards()
	expected := cards
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getAllCards() = %v; want %v", result, expected)
	}
}

func TestFillCards(t *testing.T) {
	card1 := Card{Name: "Card1"}
	card2 := Card{Name: "Card2"}
	card3 := Card{Name: "Card2"}
	filledCards := []Card{card1, card2, card3}

	FillCards(filledCards)
	result := cards
	expected := filledCards
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("fillCards(filledCards) = %v; want %v", result, expected)
	}
}

func TestRemoveAllCards(t *testing.T) {
	card1 := Card{Name: "Card3"}
	card2 := Card{Name: "Card4"}
	cards = []Card{card1, card2}

	RemoveAllCards()
	result := cards
	expected := []Card{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeAllCards() = %v; want %v", result, expected)
	}

}

func TestAddCard(t *testing.T) {
	card1 := Card{Name: "Card1"}

	AddCard(card1)
	result := cards
	expected := []Card{card1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addCard(card) = %v; want %v", result, expected)
	}
}
