package main

import (
	"fmt"

	"github.com/dmitryikh/leaves"
)

func main() {
	fmt.Println("Attempting to load model...")

	// 1. Read model
	useTransformation := false
	model, err := leaves.XGEnsembleFromFile("./model.bin", useTransformation)
	if err != nil {
		panic(err)
	}

	fmt.Println("Model loaded successfully!")
	fmt.Printf("Name: %s\n", model.Name())
	fmt.Printf("NFeatures: %d\n", model.NFeatures())
	fmt.Printf("NOutputGroups: %d\n", model.NOutputGroups())
	//fmt.Printf("NEstimators: %d\n", model.NEstimators())

	// 2. Do predictions!
	fvals := []float64{802, 1111}
	//p := model.PredictSingle(fvals, 0)
	predictions := make([]float64, 1)
	//p := model.Predict(fvals, 0, predictions)
	p := model.PredictDense(fvals, 1, 2, predictions, 0, 1)
	fmt.Printf("Prediction for %v: %f\n", fvals, p)
}
