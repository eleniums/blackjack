# blackjack
Text-based blackjack game built with Go.

## Run
To play the game:
```
go run ./cmd/game/main.go
```

## Rules
Rules used for this version of blackjack:
- Dealer hits soft 17
- Natural blackjacks pay 3:2
- Late surrender is allowed
- Double down is allowed
- Multiple splits are allowed
- Insurance is not available

## TODO
- [x] Implement double down
- [x] Make sure double down can only be done with original two cards
- [x] Only show available actions if they are possible, like double or split
- [x] Check for natural blackjacks and pay 3:2 (1.5x)
- [x] Change money to float
- [x] Implement split
- [ ] Implement late surrender
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

Resources:
- https://wizardofodds.com/games/blackjack/strategy/4-decks
- https://www.blackjackapprenticeship.com/blackjack-strategy-charts

## Links
- https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
- https://bicyclecards.com/how-to-play/blackjack
- https://www.888casino.com/blog/blackjack-strategy-guide/how-to-play-blackjack