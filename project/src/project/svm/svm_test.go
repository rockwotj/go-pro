package svm
import (
//        "svm"
//        "fmt"
        "testing"
)



func TestFirst(t *testing.T) {
        data := [][]float64{{0.0,0.0},{0.0,1.0},{1.0,0.0},{1.0,1.0}}
	labels := []float64{-1,1,1,-1}
	alpha, labels2 := train(data,labels)
	
	sample1 := []float64{0.0,0.0}
	class := predict(sample1, data, labels2, alpha)
	if class != -1 {
		t.Errorf("Incorrect class, was %.0f expected -1", class)
	}


	sample2 := []float64{1.0,0.0}
	class = predict(sample2, data, labels2, alpha)
	if class != 1 {
		t.Errorf("Incorrect class, was %.0f expected 1", class)
	}

	sample3 := []float64{0.0,1.0}
	class = predict(sample3, data, labels2, alpha)
	if class != 1 {
		t.Errorf("Incorrect class, was %.0f expected 1", class)
	}

	sample4 := []float64{1.0,1.0}
	class = predict(sample4, data, labels2, alpha)
	if class != -1 {
		t.Errorf("Incorrect class, was %.0f expected -1", class)
	}
}
