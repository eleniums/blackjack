import sys
import xgboost as xgb
import numpy as np

bst = xgb.XGBClassifier()
bst.load_model(sys.argv[1])

data = np.matrix([int(sys.argv[2]), int(sys.argv[3])])
result = bst.predict(data)
print(result[0])
