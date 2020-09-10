package game

import (
	"reflect"
	"testing"
)

func Test_mergeElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				arr: []int{2, 0, 0, 0},
			},
			want: []int{2, 0, 0, 0},
		},
		{
			name: "two",
			args: args{
				arr: []int{2, 2, 0, 0},
			},
			want:  []int{4, 0, 0, 0},
		},
		{
			name: "three",
			args: args{
				arr:  []int{4, 4, 2, 0},
			},
			want: []int{8, 2, 0, 0},
		},
		{
			name: "four",
			args: args{
				arr: []int{4, 4, 2, 2},
			},
			want: []int{8, 4, 0, 0},
		},
		{
			name: "five",
			args: args{
				arr: []int{4, 4, 4, 0},
			},
			want: []int{8, 4, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
