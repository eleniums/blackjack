#!/bin/bash
set -e

./scripts/run.sh -standard-ai -random-ai -machine-ai -machine-ai-model ./machine/model.bin -num-players 0 -num-rounds 1000 $@