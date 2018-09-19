# GoPMML
Go API for read Predictive Model Markup Language (PMML).

Currently support : Logistic Regression

next implement : Naive Bayes

Contact me at alibaltschun@gmail.com

## Installation
	go get github.com/alibaltschun/GoPMML

## Usage
	// load model
	lr,err := NewLogisticRegression("./model/logistic_regression.xml")
	if err != nil {
		panic(err)
	}
	
	// create feature
	features := map[string]float64{}
	features["x0"] = 0.1
	features["x1"] = 0.1
	features["x2"] = 0.1
	features["x3"] = 0.1

	// check empty numeric prediction array
	label, confidence, err := lr.Pred(features,true)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

## Contributing
Bug reports and pull request are wellcome

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).