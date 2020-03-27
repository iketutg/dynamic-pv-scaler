package utils

import "testing"

func TestCalculateUpdatedSize(t *testing.T) {
	type args struct {
		value      int
		percentage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCase1",
			args:args{11,10},
			want:12,
		},
		{
			name: "testCase2",
			args:args{05,30},
			want:6,
		},
		{
			name: "testCase3",
			args:args{1007,300},
			want:4028,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateUpdatedSize(tt.args.value, tt.args.percentage); got != tt.want {
				t.Errorf("CalculateUpdatedSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
