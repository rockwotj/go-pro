package svm
import (
        "svm"
        "fmt"
        "testing"
)



func TestFirst(t *testing.T) {
	data = [[0.0,0.0], [0.0,1.0],[1.0,0.0],[1.0,1.0]]
	labels = [-1,1,1,-1]
	svm.train(data,labels)
}
