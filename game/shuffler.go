package game

type Shuffler struct {
	deck Deck
}

func NewShuffler() Shuffler {
	return Shuffler{}
}

func (s Shuffler) Add(cards ...Card) {
	// TODO: shuffle cards into the internal deck
}

func (s Shuffler) Deal() Card {
	// TODO: implement deal. Might want to switch to a pointer so we can return nil
	return NewCard(1, 1)
}
