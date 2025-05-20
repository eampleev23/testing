package multiply

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "1 * 2", values: []int{1, 2}, want: 2},
		{name: "1 * 2 * 3", values: []int{1, 2, 3}, want: 6},
		{name: "1 * 2 * -4", values: []int{1, 2, -4}, want: -8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, Multiply(test.values...), test.want)
		})
	}
}
