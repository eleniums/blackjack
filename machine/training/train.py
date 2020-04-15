#!/usr/bin/env python3

"""
Train a new model with the provided dataset.
Usage: train.py <data_csv> <model_file>
"""

import sys
import numpy as np
import pandas as pd
import xgboost as xgb

print("Loading dataset: {}".format(sys.argv[1]))
dataset = pd.read_csv(sys.argv[1])

X = dataset.iloc[:, 1:3].values
y = dataset.iloc[:, 0].values

# classifier = xgb.XGBClassifier(max_depth=5,  # (Default: 6) Maximum depth of a tree. Increasing this value makes the model more complex and likely to be overfit. 0 indicates no limit.
#                                eta=.2,  # (Default: 0.3) Step size shrinkage used in updates to prevent overfitting. After each boosting step, you can directly get the weights of new features. The eta parameter actually shrinks the feature weights to make the boosting process more conservative.
#                                gamma=4,  # (Default: 0) Minimum loss reduction required to make a further partition on a leaf node of the tree. The larger, the more conservative the algorithm is.
#                                min_child_weight=6,  # (Default: 1) Minimum sum of instance weight (hessian) needed in a child. If the tree partition step results in a leaf node with the sum of instance weight less than min_child_weight, the building process gives up further partitioning. In linear regression models, this simply corresponds to a minimum number of instances needed in each node. The larger the algorithm, the more conservative it is.
#                                silent=0,  # (Default: 0) 0 means print running messages, 1 means silent mode.
#                                objective="multi:softmax",  # (Default: reg:squarederror) Specifies the learning task and the corresponding learning objective. Examples: reg:logistic, multi:softmax, reg:squarederror. For a full list of valid inputs, refer to XGBoost Parameters (https://github.com/dmlc/xgboost/blob/master/doc/parameter.rst).
#                                num_class=12,  # The number of classes. Required if objective is set to multi:softmax or multi:softprob.
#                                num_round=10)  # The number of rounds to run the training. Required.

classifier = xgb.XGBClassifier(
    objective="multi:softmax",
    num_class=12,
)

print("Starting training...")
classifier.fit(X, y)
print("Training finished")

print("Quick prediction test:")
data = np.matrix([803, 1111])
print(data)
result = classifier.predict(data)
print(result)

print("Saving model...")
classifier.save_model(sys.argv[2])
print("Model saved: {}".format(sys.argv[2]))
