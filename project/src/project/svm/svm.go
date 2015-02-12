package svm

// algorithm comes from http://math.unt.edu/~hsp0009/smo.pdf

import (
//	"fmt"
	"math"
	"math/rand"
	"time"
)


func train(data [][]float64, labels []float64) ([]float64, []float64){
	rand.Seed( time.Now().UTC().UnixNano()) // seend random number generator

	C := 1.0 // regularization parameter, lower value -> more regularization
	tol := 0.0001 // numerical tolerance (shouldn't be changed)
	max_iter := 10000 // maximum number of iterations
	num_passes := 10 // number of without any change before we decide its converged

	//kernel := "linear" // we can add "rbf" later, if we want to.
	
	N := len(data) // get length of data

	alpha := make([]float64, N) // should make an array of length N of all floats = 0
	b := 0.0
	passes := 0
	iter := 0

	for passes < num_passes && iter < max_iter { // go until converged or max iter
		alpha_change := 0 // for testing convergence

		for i:=0; i < N; i++ { // go through all data

			margin := b
			for k:=0; k < N; k++{
				margin += alpha[k]*labels[k]*kernel(data[i], data[k])
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
					margin2 += alpha[k]*labels[i]*kernel(data[j],data[k])
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
				eta := 2 * kernel(data[i],data[j]) - kernel(data[i],data[i]) - kernel(data[j],data[j])
				if eta > 0 {
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
				b1 := b - Ei - labels[i]*(newai-ai)*kernel(data[i],data[i]) -
					labels[j]*(newaj-aj)*kernel(data[i],data[j])
				b2 := b - Ej - labels[i]*(newai-ai)*kernel(data[i],data[j]) -
					labels[j]*(newaj-aj)*kernel(data[j],data[j])

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
	return alpha, labels
}


// simple linear kernel (can switch to a more complex 'rbf' kernel later, if needed)
func kernel(a []float64, b []float64) float64 {
	s := 0.0
	for i:=0; i<len(a); i++ {
		s += a[i]*b[i]
	}
	return s
}
