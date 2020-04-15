package machine

import (
	"github.com/dmitryikh/leaves"
	"github.com/eleniums/blackjack/game"
)

var model *leaves.Ensemble

func init() {
	var err error
	model, err = leaves.XGEnsembleFromFile("./machine/model.bin", false)
	if err != nil {
		panic(err)
	}
}

// Predict will feed a dealer hand and player hand into a model and return the resulting label.
func Predict(dealer *game.Hand, player *game.Hand, modelFile string, predictScript string) Label {
	d := float64(ConvertHand(FormatHand(dealer)))
	p := float64(ConvertHand(FormatHand(player)))

	predictions := make([]float64, model.NOutputGroups())
	fvals := []float64{d, p}
	err := model.Predict(fvals, 0, predictions)
	if err != nil {
		panic(err)
	}

	return Label(max(predictions))
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
