package calc

import "testing"

func TestCreateGraph(t *testing.T) {
	CreateGraph(quad)
}

func quad(x float64) float64 {
	return x*x - 2*x
}
