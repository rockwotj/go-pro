package svm

// algorithm comes from http://math.unt.edu/~hsp0009/smo.pdf

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var maxVals = [6]float64{0,0,0,0,0,0}
var minVals = [6]float64{0,0,0,0,0,0}
var f_data [][]float64
var labels []float64
var alpha []float64
var N = 0

func Normalize(data []float64) []float64 {
	for selectedType := 0; selectedType < 6; selectedType++ {	
		for feat := selectedType; feat <len(data); feat += 6 {
			data[feat] = (data[feat] - minVals[selectedType]) / (maxVals[selectedType]-minVals[selectedType])
		} 
	}
	return data
}

func NormalizeAll(data [][]float64) [][]float64 {
	for selectedType := 0; selectedType < 6; selectedType++ {
		for i := 0; i < len(data); i++ {
			for feat := selectedType; feat <len(data[0]); feat += 6 {
				if data[i][feat] < minVals[selectedType] {
					minVals[selectedType] = data[i][feat]
				}
				if data[i][feat] > maxVals[selectedType] {
					maxVals[selectedType] = data[i][feat]
				}
			} 
		}
	}

	for selectedType := 0; selectedType < 6; selectedType++ {
		for i := 0; i < len(data); i++ {
			for feat := selectedType; feat <len(data[0]); feat += 6 {
				data[i][feat] = (data[i][feat] - minVals[selectedType]) / (maxVals[selectedType]-minVals[selectedType])
			} 
		}
	}

	return data
}


func Predict(sample []float64) float64 {
	class := 0.0
	for j:=0; j<N; j++ {
//		fmt.Println("Kernel", kernel(sample, f_data[j], true))
		class += alpha[j]*labels[j]*kernel(sample, f_data[j], false)
	}
	fmt.Println("%.2f", class)
	if class > 0 {
		return 1
	}
	return -1
}


func Train(data [][]float64, label []float64) ([]float64, []float64){
	rand.Seed( time.Now().UTC().UnixNano()) // seend random number generator
	labels = label
	f_data = data

	C := 2.05 // regularization parameter, lower value -> more regularization
	tol := 0.0001 // numerical tolerance (shouldn't be changed)
	max_iter := 10000 // maximum number of iterations
	num_passes := 10 // number of without any change before we decide its converged

	//kernel := "linear" // we can add "rbf" later, if we want to.
	
	N = len(data) // get length of data

	alpha = make([]float64, N) // should make an array of length N of all floats = 0
	b := 0.0
	passes := 0
	iter := 0

	for passes < num_passes && iter < max_iter { // go until converged or max iter
		alpha_change := 0 // for testing convergence

		for i:=0; i < N; i++ { // go through all data

			margin := b
			for k:=0; k < N; k++{
				margin += alpha[k]*labels[k]*kernel(data[i], data[k], false)
			}
			Ei := margin - labels[i] // eq 2


			if (labels[i]*Ei < -tol && alpha[i] < C) ||
				(labels[i]*Ei > tol && alpha[i] > 0) {
				
				// update alpha_i
				j := i
				for j == i {
					j = rand.Intn(N) // random number in [0,N)
				}
				margin2 := b
				for k:=0; k<N; k++{
					margin2 += alpha[k]*labels[k]*kernel(data[j],data[k], false)
				}
				Ej := margin2 - labels[j]
				
				// calc L and H (lower and upper bounds) by eq 10/11
				// and save old alpha_i,j
				ai := alpha[i]
				aj := alpha[j]
				L := 0.0
				H := C
				
				if labels[i] == labels[j] {
					L = math.Max(0, ai+aj-C)
					H = math.Min(C, ai+aj)
				} else {
					L = math.Max(0, aj-ai)
					H = math.Min(C, C+aj-ai)
				}

				if math.Abs(L-H) < 0.0001 {
					continue
				}
				//compute eta by eq 14
				eta := 2 * kernel(data[i],data[j],false) - kernel(data[i],data[i], false) - kernel(data[j],data[j], false)
				if eta >= 0 {
					continue
				}

				newaj := aj - labels[j]*(Ei-Ej) / eta // eq 12,15
				if newaj > H {
					newaj = H
				}
				if newaj < L {
					newaj = L
				}

				if math.Abs(aj-newaj) < 0.0001 {
					continue
				}
				alpha[j] = newaj

				// compute new alpha_i by eq 16
				newai := ai + labels[i]*labels[j]*(aj-newaj)
				alpha[i] = newai

				// compute b1, b2 with eq 17, 18
				b1 := b - Ei - labels[i]*(newai-ai)*kernel(data[i],data[i], false) -
					labels[j]*(newaj-aj)*kernel(data[i],data[j], false)
				b2 := b - Ej - labels[i]*(newai-ai)*kernel(data[i],data[j], false) -
					labels[j]*(newaj-aj)*kernel(data[j],data[j], false)

				// compute b by eq 19
				b = 0.5*(b1+b2)

				if newai > 0 && newai < C {
					b = b1
				}
				if newaj > 0 && newaj < C {
					b = b2
				}
				alpha_change++
				
			} // alpha_i update (if)
		} // data loop
		
		iter++
		if alpha_change == 0 {
			passes++
		} else {
			passes = 0
		}
	} // end training loop
	
	// return alphas and labels
	f_data = data
	return alpha, labels
}


func kernel(a []float64, b []float64, x bool) float64 {
	return rbfkernel(a, b, x)
}

// simple linear kernel (can switch to a more complex 'rbf' kernel later, if needed)
func linearkernel(a []float64, b []float64) float64 {
	s := 0.0
	for i:=0; i<len(a); i++ {
		s += a[i]*b[i]
	}
	return s
}


func rbfkernel(a []float64, b []float64, x bool) float64 {
	sigma := 0.75
	s := 0.0
	for i:=0; i<len(a); i++ {
		s += (a[i] - b[i]) * (a[i] - b[i])
	}
	return math.Exp(-s / (2.0 * sigma * sigma))
}
