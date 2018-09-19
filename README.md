# GoPMML
Go API for read Predictive Model Markup Language (PMML).

Currently support : Logistic Regression

next implement : Naive Bayes

Contact me at alibaltschun@gmail.com

## Installation
	go get github.com/alibaltschun/GoPMML

## Usage
	// load model
	modelXML, _ := ioutil.ReadFile(fileModel)
	var model LogisticRegression
	err := xml.Unmarshal([]byte(modelXML), &model)
	
	// create feature
	features1 := map[string]float64{}
	features1["x0"] = 0.1
	features1["x1"] = 0.1
	features1["x2"] = 0.1
	features1["x3"] = 0.1

	// check empty numeric prediction array
	label, confidence, err := model.Score(features1)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

## Contributing
Bug reports and pull request are wellcome

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).