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
- Cannot surrender after split
- Double down is allowed
- Double after split is allowed
- Multiple splits are allowed
- Insurance is not available

## Computer AIs
Computer AIs can be added to the game using command-line flags.

### Random AI
Flag: `-random-ai`

Uses a random number generator to either hit or stay on every move.

### Standard AI
Flag: `-standard-ai`

Plays using a standard blackjack strategy.

Strategy Resources:
- https://wizardofodds.com/games/blackjack/strategy/4-decks
- https://www.blackjackapprenticeship.com/blackjack-strategy-charts

## Links
- https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
- https://bicyclecards.com/how-to-play/blackjack
- https://www.888casino.com/blog/blackjack-strategy-guide/how-to-play-blackjack