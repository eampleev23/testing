package multiply

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMul(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "1 and 2", values: []int{1, 2}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, Mul(test.values...), test.want)
		})
	}
}
