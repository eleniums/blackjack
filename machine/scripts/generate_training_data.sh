#!/bin/bash
set -e

iterations=1000
output="./machine/testdata/training.csv"
converted="./machine/testdata/training_converted.csv"

echo "Running simulation to generate training data..."
./scripts/run.sh -standard-ai -random-ai -num-players 0 -num-rounds $iterations -generate-training-data -training-data-file $output $@

echo "Converting training data to an appropriate format..."
go run ./machine/training/convert.go $output $converted

echo "Training data generated: $output"
echo "Training data converted: $converted"
