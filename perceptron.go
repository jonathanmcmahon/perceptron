// An implementation of Rosenblatt's perceptron algorithm

package perceptron

// A Perceptron is a single-layer Rosenblatt perceptron.
type Perceptron struct {
	bias    bool      // use bias term
	weights []float64 // perceptron weights
}

// New returns a new perceptron with all weights initalized to zero (and, optionally, a bias term).
func New(bias bool, inputs uint) Perceptron {
	initial_value := 0.0
	var weights []float64

	// Add bias term
	if bias {
		weights = append(weights, 1)
	}

	// Initialize all weights to very small values
	for i := uint(0); i < inputs; i++ {
		weights = append(weights, initial_value)
	}
	p := Perceptron{bias, weights}
	return p
}
