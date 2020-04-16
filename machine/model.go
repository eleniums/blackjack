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

// Predict will feed a dealer hand and player hand into a model and return the resulting prediction with the highest score.
func (m *Model) Predict(dealer *game.Hand, player *game.Hand) Prediction {
	predictions := m.predictInternal(dealer, player)

	// get the index of the highest value in the array
	index := 0
	maxValue := 0.0
	for i, v := range predictions {
		if v > maxValue {
			index = i
			maxValue = v
		}
	}

	action, result := Label(index).Split()

	return Prediction{
		Action: action,
		Result: result,
		Score:  maxValue,
	}
}

// PredictAll will feed a dealer hand and player hand into a model and return all resulting predictions with their scores.
func (m *Model) PredictAll(dealer *game.Hand, player *game.Hand) []Prediction {
	predictions := m.predictInternal(dealer, player)

	results := []Prediction{}
	for i, v := range predictions {
		action, result := Label(i).Split()
		results = append(results, Prediction{
			Action: action,
			Result: result,
			Score:  v,
		})
	}

	return results
}

// predictInternal will feed a dealer hand and player hand into a model and return the resulting predictions.
func (m *Model) predictInternal(dealer *game.Hand, player *game.Hand) []float64 {
	d := float64(ConvertHand(FormatHand(dealer)))
	p := float64(ConvertHand(FormatHand(player)))

	predictions := make([]float64, m.model.NOutputGroups())
	fvals := []float64{d, p}
	err := m.model.Predict(fvals, 0, predictions)
	if err != nil {
		panic(err)
	}

	return predictions
}
