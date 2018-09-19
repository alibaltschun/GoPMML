package gopmml

import (
	"fmt"
	"testing"
)

func Test_parsinglogreg(t *testing.T) {
	lr, err := NewLogisticRegression("./model/logistic_regression.xml")
	if err != nil {
		t.Error(err.Error())
		t.Fail()
		return
	}

	features1 := map[string]float64{}
	features1["x0"] = 0.1
	features1["x1"] = 0.1
	features1["x2"] = 0.1
	features1["x3"] = 0.1

	// check empty numeric prediction array
	label, confidence, err := lr.Pred(features1, true)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

	lr.SetupNumbericPredictorMap()

	// check normal case
	label, confidence, err = lr.Pred(features1, true)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

	// check normal case without normalization
	label, confidence, err = lr.Pred(features1, false)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

	// check empty feature
	features0 := map[string]float64{}
	label, confidence, err = lr.Pred(features0, true)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

	// check empty feature without normalization
	label, confidence, err = lr.Pred(features0, false)
	fmt.Println(label)
	fmt.Println(confidence)
	fmt.Println(err)

	// check softnormalization with empty feature
	prob, err := SoftmaxNormalizationMethods(nil)
	fmt.Println(prob)
	fmt.Println(err)
}
