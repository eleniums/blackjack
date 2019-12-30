# blackjack
Text-based blackjack game built with Go.

## Run
To play the game:
```
go run ./cmd/game/main.go
```

## TODO
- [x] Implement double down
- [ ] Make sure double down can only be done with original two cards
- [ ] Only show available actions if they are possible, like double or split
- [ ] Implement split
- [ ] Implement standard AI
- [ ] Add code to create training data (show hand and potential outcomes, like HITLOSS, HITWIN, STAYLOSS, STAYWIN, DOUBLEWIN, DOUBLELOSS, SPLIT)
- [ ] Create machine learning AI

## Computer AIs
Computer AIs can be added to the game using command-line flags.

### Random AI
Flag: `-random-ai`

Uses a random number generator to either hit or stay on every move.

### Standard AI
Flag: `-standard-ai`

Plays using a standard blackjack strategy.

## Links
- https://en.wikipedia.org/wiki/Playing_cards_in_Unicode