package stack

import "testing"

func Test_stack(t *testing.T) {
	type args struct {
		array []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Stack Test",
			args: args{
				array: []string{"a", "b", "c", "d", "e"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack(tt.args.array)
		})
	}
}
