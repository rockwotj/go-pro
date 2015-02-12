package svm
import (
//        "svm"
//        "fmt"
        "testing"
)



func TestFirst(t *testing.T) {
        data := [][]float64{{0.0,0.0},{0.0,1.0},{1.0,0.0},{1.0,1.0}}
	labels := []float64{-1,1,1,-1}
	train(data,labels)
}
