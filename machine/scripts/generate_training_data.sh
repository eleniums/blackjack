#!/bin/bash
set -e

output="./machine/testdata/training.csv"

./scripts/run.sh -standard-ai -random-ai -num-players 0 -num-rounds 1000 -generate-training-data -training-data-file $output $@

echo "Training data generated: $output"
