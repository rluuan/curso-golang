package matematica

import (
	"testing"
)

const errorDefault = "Expected value %v, output value equal %v"

func TestWeightedAverage(t *testing.T) {
	t.Parallel()
	const expectedValue = 7.10

	getValueWeightedAverage := WeightedAverage(5.66, 7.44, 8.21)

	if getValueWeightedAverage != expectedValue {
		t.Error(errorDefault, expectedValue, getValueWeightedAverage)
	}
}
