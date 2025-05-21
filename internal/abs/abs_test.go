package abs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAbs более громоздкая реализация стандартной библиотекой.
func TestAbs(t *testing.T) {

	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "-3", value: -3, want: 3},
		{name: "3", value: 3, want: 3},
		{name: "-2.000001", value: -2.000001, want: 2.000001},
		{name: "-0.000000003", value: -0.000000003, want: 0.000000003},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Abs(test.value); got != test.want {
				t.Errorf("%s fails.. Abs() = %v, want %v", test.name, got, test.want)
			}
		})
	}

}

// TestAbs2 реализация теста с использованием сторонней библиотеки.
func TestAbs2(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "negative value", value: -3.001, want: 3.001},
		{name: "small value", value: -0.0000001, want: 0.0000001},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.value))
		})
	}
}
