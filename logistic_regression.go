package gopmml

import (
	"encoding/xml"
	"io/ioutil"
)

// ===========================================================================================
// ===========================================================================================
// ===========================================================================================
// ===========================================================================================

type PMMLLR struct {
	// struct xml:PMML
	LogisticRegression LogisticRegression `xml:"RegressionModel"`
}

type LogisticRegression struct {
	// struct xml:PMML>RegressionModel
	NormalizationMethod string            `xml:"normalizationMethod,attr"`
	Fields              []MiningField     `xml:"MiningSchema>MiningField"`
	RegressionTable     []RegressionTable `xml:"RegressionTable"`
}

type MiningField struct {
	// struct xml:PMML>RegressionModel>MiningSchema>MiningField
	Name string `xml:"name,attr"`
}

type RegressionTable struct {
	// struct xml:PMML>RegressionModel>RegressionTable
	Intercept           float64            `xml:"intercept,attr"`
	TargetCategory      string             `xml:"targetCategory,attr"`
	NumericPredictor    []NumericPredictor `xml:"NumericPredictor"`
	NumericPredictorMap *map[string]float64
}

type NumericPredictor struct {
	// struct xml:PMML>RegressionModel>RegressionTable>NumbericPredictor
	Name        string  `xml:"name,attr"`
	Coefficient float64 `xml:"coefficient,attr"`
}

// ===========================================================================================
// ===========================================================================================
// ===========================================================================================
// ===========================================================================================

// method for convert pmml file into golang object
// input  : Logistic Regression PMML file path
// output : Golang Logistic Regression model
func NewLogisticRegression(fileModel string) (LogisticRegression, error) {
	model := LogisticRegression{}

	// get binary data of pmml file
	modelXML, err := ioutil.ReadFile(fileModel)
	if err != nil {
		return model, err
	}

	// transform binary xml into model object
	err = xml.Unmarshal(modelXML, &model)
	if err != nil {
		return model, err
	}

	// initial model to extract numberic predictor
	// into map variable for ease access
	model.SetupNumbericPredictorMap()

	// return golang logistic regression object
	return model, nil
}

// method for score test data
// input : 	independent variable with map["var name"]value
//			normalize with boolean type
//				true (default)  -> using normalization
//				false			-> without normalization
// return : -label with string type
//			-confident/prob with map type
//			-errors
func (lr *LogisticRegression) Pred(features map[string]float64, normalize bool) (string, map[string]float64, error) {

	// calculate confident value using log reg function
	confident := lr.RegressionFunction(features)

	// return label without normalization
	if !normalize {
		return ArgMax(confident), confident, nil
	}

	// calculate confident value with normalization method
	var normMethod NormalizationMethodMap
	if lr.NormalizationMethod != "" {

		// check if this library doesnt support normalization used by model
		if _, ok := NormalizationMethodMaps[lr.NormalizationMethod]; !ok {
			return "", nil, NormalizationMethodNotImplemented

			// get normalization method from utils
		} else {
			normMethod = NormalizationMethodMaps[lr.NormalizationMethod]
		}
	}

	// calculate probability each class
	prob, err := normMethod(confident)
	if err != nil {
		return "", nil, err
	}

	// return label with normalization
	return ArgMax(prob), prob, nil
}

// create map for containing numeric predictor / weight
func (lr *LogisticRegression) SetupNumbericPredictorMap() {

	// get all regression table
	for i, rt := range lr.RegressionTable {
		m := make(map[string]float64)
		for _, np := range rt.NumericPredictor {

			// check if the model not used nlp variable
			if getSubstringInsideParentheses(np.Name) != "" {
				m[np.Name] = np.Coefficient

				// model used nlp variable
			} else {
				m[getSubstringInsideParentheses(np.Name)] = np.Coefficient
			}
		}

		// update numeric predictor map in regression table
		lr.RegressionTable[i].NumericPredictorMap = &m
	}
}

// method for calculate feature using logistic regression
// function for countinous independent variable
func (lr *LogisticRegression) RegressionFunction(features map[string]float64) map[string]float64 {
	confidence := map[string]float64{}

	// get all regressionTable for calculate confident
	// features of every label
	for _, regressionTable := range lr.RegressionTable {
		var intercept float64

		// get intercept of regression table
		intercept = regressionTable.Intercept

		// check if the numeric predictor map is not empty
		if regressionTable.NumericPredictorMap != nil {

			// get pointer of numeric predictor map
			m := *regressionTable.NumericPredictorMap
			sum := 0.0

			// calculate the multification of coefficient
			// with value of sub features
			for k, v := range features {
				if c, ok := m[k]; ok {
					sum += v * c
				}
			}

			// append all confident value of each class
			confidence[regressionTable.TargetCategory] = intercept + sum
		}
	}

	// return confidence value
	return confidence
}
