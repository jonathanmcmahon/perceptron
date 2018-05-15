package perceptron

import (
	"fmt"
	"testing"
)

func TestNewPerceptronWithoutBias(t *testing.T) {
	var n_inputs uint
	n_inputs = 10
	p := New(false, n_inputs)

	fmt.Println(p)

	if uint(len(p.weights)) != n_inputs {
		t.Error("Number of weights is ", len(p.weights), "; number of inputs is", n_inputs)
	}

	for _, v := range p.weights {
		if v != 0.0 {
			t.Error("Weights were not initialized to zero.")
		}
	}

}

func TestNewPerceptronWithBias(t *testing.T) {
	var n_inputs uint
	n_inputs = 10
	p := New(true, n_inputs)

	fmt.Println(p)

	if uint(len(p.weights)) != n_inputs+1 {
		t.Error("Number of weights is ", len(p.weights)-1, "; number of inputs is", n_inputs)
	}

	for i, v := range p.weights {
		if i == 0 {
			if v != 1.0 {
				t.Error("Bias term was not set properly; bias was", v)
			}
		} else {
			if v != 0.0 {
				t.Error("Weights were not initialized to zero.")
			}
		}
	}
}
