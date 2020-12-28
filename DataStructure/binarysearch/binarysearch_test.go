package binarysearch

import "testing"

func Test_binarysearch(t *testing.T) {
	type args struct {
		array     []int
		targetNum int
		start     int
		end       int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "binary search",
			args: args{
				array:     []int{1, 2, 3, 4, 5, 6},
				targetNum: 2,
				start:     0,
				end:       5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarysearch(tt.args.array, tt.args.targetNum, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("binarysearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
