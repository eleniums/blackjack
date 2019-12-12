package game

type Hand []Card

func NewHand(cards ...Card) Hand {
	var hand Hand
	for _, v := range cards {
		hand = append(hand, v)
	}
	return hand
}
