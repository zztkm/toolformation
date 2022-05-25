package toolformation

import "testing"

func TestCheck(t *testing.T) {
	var tests = []struct {
		name string
		args string
		want int
	}{
		{"normal", "which", 0},
		{"abnormal", "dfafaff0921jfad", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
