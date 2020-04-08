#!/bin/bash
set -e

dataset="./machine/testdata/training_converted.csv"
model="./machine/testdata/model.bin"

./machine/training/train.py $dataset $model
