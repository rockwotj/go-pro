package svm
import (
        "testing"
	"project/imageProcessor"
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


func TestSecond(t *testing.T){
	sunsets := imageProcessor.ProcessDirectory("../../../TrainSunset/*.jpg")
	nonsunsets := imageProcessor.ProcessDirectory("../../../TrainNonsunsets/*.jpg")
	if len(sunsets) < 10 {
		t.Errorf("Incorrect sunset length", len(sunsets))
	}
	if len(nonsunsets) < 10 {
		t.Errorf("Incorrect nonsunset length", len(nonsunsets))
	}

	labelsSunset := make([]float64,len(sunsets))
	for i :=0; i < len(sunsets); i++{
		labelsSunset[i] = 1
	}
	labelsNonsunset := make([]float64,len(nonsunsets))
	for i :=0; i < len(nonsunsets); i++{
		labelsNonsunset[i] = -1
	}
	labels := append(labelsSunset, labelsNonsunset...)
	data := append(sunsets, nonsunsets...)
	alpha, labels2 := train(data,labels)

	ts := imageProcessor.ProcessDirectory("../../../TestSunset/*.jpg")
	tns := imageProcessor.ProcessDirectory("../../../TestNonsunsets/*.jpg")

	class := predict(ts[0], data, labels2, alpha)
	if class != 1 {
		t.Errorf("Incorrect class, was %.0f expected 1", class)
	}
	class = predict(ts[50], data, labels2, alpha)
	if class != 1 {
		t.Errorf("Incorrect class, was %.0f expected 1", class)
	}

	class = predict(tns[0], data, labels2, alpha)
	if class != -1 {
		t.Errorf("Incorrect class, was %.0f expected -1", class)
	}
	class = predict(tns[0], data, labels2, alpha)
	if class != -1 {
		t.Errorf("Incorrect class, was %.0f expected -1", class)
	}	
}
