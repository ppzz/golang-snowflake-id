package main

import (
	"reflect"
	"testing"
)

func Test_uniq(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			name: "case1",
			list: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "case2",
			list: []int{0, 1, 1, 2, 2, 2, 3, 4, 5, 6, 6, 6, 7, 7, 7, 7, 8, 9, 10},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "case3",
			list: []int{0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniq(tt.list)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
