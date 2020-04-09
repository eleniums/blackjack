#!/bin/bash
set -e

./scripts/run.sh -standard-ai -random-ai -machine-ai -machine-ai-model ./machine/model.bin -machine-ai-predict-script ./machine/predict.py -num-players 0 -num-rounds 5 $@