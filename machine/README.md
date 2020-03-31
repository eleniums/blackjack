# Machine Learning
Creating an AI to play Blackjack.

## Creating training data
Flags:
- `-generate-training-data`: Add to generate training data.
- `-training-data-file`: Add to specify a training data output file.

This script will generate a large set of training data using AI players:
```
./machine/scripts/generate_training_data.sh
```

## Convert training data
Training data needs to be in a numerical format:
```
go run ./machine/training/convert.go ./machine/testdata/training.csv ./machine/testdata/output.csv
```

## Training model with AWS SageMaker
Dashboard:
https://us-west-2.console.aws.amazon.com/sagemaker/home

Getting Started:
https://docs.aws.amazon.com/sagemaker/latest/dg/gs-console.html

Data Formats:
https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-training.html

Built-in Algorithms:
https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html

Training Example:
https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-train-model.html

XGBoost Hyperparameters:
https://docs.aws.amazon.com/sagemaker/latest/dg/xgboost_hyperparameters.html
