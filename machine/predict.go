package machine

import (
	"github.com/dmitryikh/leaves"
	"github.com/eleniums/blackjack/game"
)

// Model can make predictions from a model.
type Model struct {
	model *leaves.Ensemble
}

// NewModel will load the specified model file.
func NewModel(modelFile string) *Model {
	model, err := leaves.XGEnsembleFromFile(modelFile, false)
	if err != nil {
		panic(err)
	}

	return &Model{
		model: model,
	}
}

// Predict will feed a dealer hand and player hand into a model and return the resulting label.
func (m *Model) Predict(dealer *game.Hand, player *game.Hand) Label {
	d := float64(ConvertHand(FormatHand(dealer)))
	p := float64(ConvertHand(FormatHand(player)))

	predictions := make([]float64, m.model.NOutputGroups())
	fvals := []float64{d, p}
	err := m.model.Predict(fvals, 0, predictions)
	if err != nil {
		panic(err)
	}

	n := max(predictions)

	return Label(n)
}

// max will return the index of the highest value in the array.
func max(values []float64) int {
	index := 0
	maxValue := 0.0
	for i, v := range values {
		if v > maxValue {
			index = i
			maxValue = v
		}
	}
	return index
}
