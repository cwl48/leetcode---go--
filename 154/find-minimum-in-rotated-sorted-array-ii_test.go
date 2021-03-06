package findminimuminr154

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{2, 2, 2, 0, 1}}, 0},
		{"test", args{[]int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"test", args{[]int{1, 3, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
