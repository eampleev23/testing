package sum

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "sum 1 and 2", values: []int{1, 2}, want: 3},
		{name: "only 1", values: []int{1}, want: 1},
		{name: "with negative values", values: []int{-1, -2, 3}, want: 0},
		{name: "with negative zero", values: []int{-0, 3}, want: 3},
		{name: "a lot of values", values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, want: 189},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if sum := Sum(test.values...); sum != test.want {
				t.Errorf("%s fails.. Sum() = %d, want %d", test.name, sum, test.want)
			}
		})
	}
}
