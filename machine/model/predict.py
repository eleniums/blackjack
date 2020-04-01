import xgboost as xgb
import numpy as np

bst = xgb.XGBClassifier()
bst.load_model('model.bin')

data = np.matrix([803, 1111])
print(data)
result = bst.predict(data)
print(result)
