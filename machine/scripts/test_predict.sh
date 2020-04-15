#!/bin/bash
set -e

model="./machine/testdata/model.bin"
dealer=803
player=1111

echo "$dealer $player"
./machine/predict.py $model $dealer $player
