package gopmml

import (
	"errors"
	"math"
	"strings"
)

type NormalizationMethodMap func(map[string]float64) (map[string]float64, error)

var NormalizationMethodMaps map[string]NormalizationMethodMap
var NormalizationMethodNotImplemented = errors.New("Normalization Method Not Implemented Yet")

func init() {
	NormalizationMethodMaps = map[string]NormalizationMethodMap{}
	NormalizationMethodMaps["softmax"] = SoftmaxNormalizationMethods
}

// method for check if the idependent variable is nlp
func getSubstringInsideParentheses(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s[i:], ")")
		if j >= 0 {
			return s[i+1 : j+i]
		}
	}
	return ""
}

// function for compute confidence value
// into probability using softMax function
// input  : map of confidence value with float64 type
// output : map of probability each class with float64 type
func SoftmaxNormalizationMethods(confidence map[string]float64) (map[string]float64, error) {
	if confidence != nil {
		result := map[string]float64{}
		tempExp := []float64{}
		for _, v := range confidence {
			tempExp = append(tempExp, math.Exp(v))
		}
		sum := 0.0
		for _, j := range tempExp {
			sum += j
		}

		i := 0
		for k, _ := range confidence {
			result[k] = tempExp[i]
			i += 1
		}
		return result, nil
	}
	return nil, errors.New("feature is empty")
}

// method for get key with search max value in map
func ArgMax(feature map[string]float64) string {
	result := ""
	max := -999.999
	for k, v := range feature {
		if result != "" {
			if max < v {
				result = k
				max = v
			}
		} else {
			result = k
			max = v
		}
	}
	return result
}
