package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	kind int
	suit int
}

func DescribeCard(c *Card) string {
	kind := ""
	suit := ""

	if c.kind == 1 {
		kind = "Ace"
	} else if c.kind == 11 {
		kind = "Jack"
	} else if c.kind == 12 {
		kind = "Queen"
	} else if c.kind == 13 {
		kind = "King"
	} else {
		kind = strconv.Itoa(c.kind)
	}

	if c.suit == 0 {
		suit = "Hearts"
	} else if c.suit == 1 {
		suit = "Diamonds"
	} else if c.suit == 2 {
		suit = "Clubs"
	} else if c.suit == 3 {
		suit = "Spades"
	}

	return kind + " of " + suit
}

func MakeDeck() []Card {
	var deck []Card

	for kind := 1; kind <= 13; kind++ {
		for suit := 0; suit < 4; suit++ {
			card := Card{kind, suit}
			deck = append(deck, card)
		}
	}

	return deck
}

func PickRandomCard(deck []Card) ([]Card, Card) {
	i := rand.Intn(53)
	card := deck[i]

	deck = append(deck[:i], deck[i+1:]...)

	return deck, card
}

func Shuffle(deck []Card) []Card {

	shuffledDeck := make([]Card, len(deck))
	perm := rand.Perm(len(deck))

	for i := 0; i < len(perm); i++ {
		shuffledDeck[i] = deck[perm[i]]
	}

	return shuffledDeck
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	deck := MakeDeck()
	deck, c := PickRandomCard(deck)
	fmt.Println(DescribeCard(&c))

	deck = Shuffle(deck)
	for _, c := range deck {
		fmt.Println(DescribeCard(&c))
	}
}
