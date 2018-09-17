package main

import "testing"

func TestPrintNumbers(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ten", args{8}, 8},
		{"0", args{0}, 0},
		{"negative", args{-5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintNumbers(tt.args.m); got != tt.want {
				t.Errorf("PrintNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
