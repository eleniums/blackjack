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

## Results
Not too shabby for a first stab. Marvin is the player using the trained model.
```
Larry (*ai.Random)
  Win: 313 (%31.3) | Loss: 641 (%64.1) | Tie: 46 (%4.6) | $-4497.50
Joe (*ai.Standard)
  Win: 439 (%42.8) | Loss: 507 (%49.5) | Tie: 79 (%7.7) | $-65.00
Marvin (*ai.Machine)
  Win: 429 (%41.6) | Loss: 534 (%51.8) | Tie: 68 (%6.6) | $-1100.00
```
