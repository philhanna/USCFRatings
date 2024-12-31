package uscfratings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgeBasedRating(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{age: 10, want: 500},
		{age: 1, want: 100},
		{age: 26, want: 1300},
		{age: 70, want: 1300},
		{age: 0, want: 1300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := AgeBasedRating(tt.age)
			assert.Equal(t, want, have)
		})
	}
}

