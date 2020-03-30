package machine

import (
	"os"

	"github.com/eleniums/blackjack/game"
)

// Recorder contains methods for recording machine learning training data.
type Recorder struct {
	data   *os.File
	record *record
}

// NewRecorder will create a new Recorder instance with an open file for storing training data.
func NewRecorder(file string) *Recorder {
	data, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return &Recorder{
		data:   data,
		record: nil,
	}
}

// Close the open file handle. This should always be called.
func (r *Recorder) Close() {
	r.data.Close()
}

// Start will begin a new record.
func (r *Recorder) Start(dealer, player *game.Hand, action game.Action) {
	if r == nil {
		return
	}

	// close out any existing record
	if r.record != nil {
		r.Write(game.ResultInvalid)
	}

	r.record = &record{
		action: action,
	}
	r.record.AddDealerHand(dealer)
	r.record.AddPlayerHand(player)
}

// Write a completed record to a file.
func (r *Recorder) Write(result game.Result) {
	if r == nil || r.record == nil {
		return
	}

	r.record.result = result
	r.record.Write(r.data)
	r.record = nil
}
