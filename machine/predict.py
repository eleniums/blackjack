#!/usr/bin/env python3

"""
Load model and retrieve prediction.
Usage: predict.py <model_file> <dealer_hand> <player_hand>
"""

import sys
import xgboost as xgb
import numpy as np

bst = xgb.XGBClassifier()
bst.load_model(sys.argv[1])

data = np.matrix([int(sys.argv[2]), int(sys.argv[3])])
result = bst.predict(data)
print(result[0])
